package mongodb

import (
	"backend/internal/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

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

	fmt.Println("Successfully connected to MongoDB")

	return err
}
