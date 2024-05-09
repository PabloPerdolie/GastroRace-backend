package productrepo

import (
	"backend/internal/models"
	"backend/internal/mongodb"
	"bytes"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func CreateProduct(ctx context.Context, product models.Product) (prod models.Product, err error) {
	id := primitive.NewObjectID()
	stream, err := mongodb.FS.OpenUploadStreamWithID(id, product.Name)
	if err != nil {
		return prod, err
	}
	defer stream.Close()
	_, err = stream.Write(product.ImageData)
	if err != nil {
		return prod, err
	}
	product.ImageId = id
	one, err := mongodb.ProductColl.InsertOne(ctx, product)
	if err != nil {
		return prod, err
	}
	log.Printf("Inserted with id = %v", one.InsertedID)
	return product, nil
}

func DeleteProduct(ctx context.Context, id primitive.ObjectID) error {
	one, err := FindOne(context.Background(), id)
	if err != nil {
		return err
	}
	imageId := one.ImageId
	filter := bson.M{"_id": id}
	res, err := mongodb.ProductColl.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("product not found")
	}

	log.Printf("Deleted %d documents", res.DeletedCount)
	if err := mongodb.FS.Delete(imageId); err != nil {
		return err
	}
	if err := mongodb.FS.Delete(imageId); err != nil {
		return err
	}
	return nil
}

func FindOne(ctx context.Context, id primitive.ObjectID) (prod models.Product, err error) {
	filter := bson.M{"_id": id}
	result := mongodb.ProductColl.FindOne(ctx, filter)
	if result.Err() != nil {
		return models.Product{}, result.Err()
	}
	if err = result.Decode(&prod); err != nil {
		return prod, err
	}

	downloadStream, err := mongodb.FS.OpenDownloadStream(prod.ImageId)
	if err != nil {
		return prod, err
	}

	_, err = downloadStream.Read(prod.ImageData)
	if err != nil {
		return prod, err
	}
	return prod, err
}

func FindAll(ctx context.Context) (prods []models.Product, err error) {
	result, err := mongodb.ProductColl.Find(ctx, bson.D{{}})
	if result.Err() != nil {
		return prods, result.Err()
	}
	if err = result.All(ctx, &prods); err != nil {
		return prods, fmt.Errorf("failed to read file data from cursor: %v", err)
	}
	log.Println(prods)
	for i, prod := range prods {
		var buf bytes.Buffer
		_, err = mongodb.FS.DownloadToStream(prod.ImageId, &buf)
		if err != nil {
			return prods, err
		}
		prods[i].ImageData = buf.Bytes()
	}

	return prods, err
}
