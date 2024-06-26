package mongodb

import (
	"backend/internal/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var DB *mongo.Database
var FS *gridfs.Bucket
var ProductColl *mongo.Collection
var UserColl *mongo.Collection
var OrderColl *mongo.Collection

func InitDB(ctx context.Context) (err error) {
	clientOptions := options.Client().ApplyURI(config.CONFIG.DB.Url)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	DB = client.Database(config.CONFIG.DB.Name)
	FS, err = gridfs.NewBucket(client.Database(config.CONFIG.DB.Name))
	if err != nil {
		return err
	}
	ProductColl = DB.Collection("products")
	UserColl = DB.Collection("users")
	OrderColl = DB.Collection("orders")

	log.Println("Successfully connected to MongoDB")

	return err
}
