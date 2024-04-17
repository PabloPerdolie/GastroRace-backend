package productrepo

import (
	"backend/internal/models"
	"backend/internal/mongodb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

var db = mongodb.DB.Collection("products")

func CreateProduct(ctx context.Context, product models.Product) error {
	id := primitive.NewObjectID()
	stream, err := mongodb.FS.OpenUploadStreamWithID(id, "image.jpg")
	if err != nil {
		return err
	}
	defer stream.Close()
	_, err = stream.Write(product.ImageData)
	if err != nil {
		return err
	}
	product.ImageId = id
	one, err := db.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	log.Printf("Inserted with id = %v", one.InsertedID)
	return nil
}

func DeleteProduct(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	res, err := db.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("product not found")
	}
	log.Printf("Deleted %d documents", res.DeletedCount)
	// todo find product to delete with ImageID from GridFS
	return nil
}

func FindOne(ctx context.Context, id primitive.ObjectID) (prod models.Product, err error) {
	filter := bson.M{"_id": id}
	result := db.FindOne(ctx, filter)
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

func FindAll(ctx context.Context, id primitive.ObjectID) (prods []models.Product, err error) {
	result, err := db.Find(ctx, bson.D{{}})
	if err != nil {
		return prods, err
	}
	if err = result.Decode(&prods); err != nil {
		return prods, err
	}
	for _, prod := range prods {
		downloadStream, err := mongodb.FS.OpenDownloadStream(prod.ImageId)
		if err != nil {
			return prods, err
		}

		_, err = downloadStream.Read(prod.ImageData)
		if err != nil {
			return prods, err
		}
	}

	return prods, err
}
