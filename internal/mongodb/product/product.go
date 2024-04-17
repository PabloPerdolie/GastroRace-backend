package product

import (
	"backend/internal/models"
	"backend/internal/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

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
	one, err := mongodb.DB.Collection("products").InsertOne(ctx, product)
	if err != nil {
		return err
	}
	log.Printf("Inserted with id = %v", one.InsertedID)
	return nil
}

func DeleteProduct(ctx context.Context, id primitive.ObjectID) error {

	return nil
}
