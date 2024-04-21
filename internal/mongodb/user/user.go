package userrepo

import (
	"backend/internal/models"
	"backend/internal/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func CreateUser(ctx context.Context, user models.User) error {
	one, err := mongodb.UserColl.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	log.Printf("Inserted with id = %v", one.InsertedID)
	return nil
}

func FindOneUser(username, password string) (user models.User, err error) {
	filter := bson.M{
		"username": username,
		"password": password,
	}
	result := mongodb.UserColl.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return user, result.Err()
	}
	if err = result.Decode(&user); err != nil {
		return user, err
	}
	return user, err
}

func FindByIdUser(ctx context.Context, id primitive.ObjectID) (user models.User, err error) {
	filter := bson.M{
		"_id": id,
	}
	result := mongodb.UserColl.FindOne(context.Background(), filter)
	if result.Err() != nil {
		return user, result.Err()
	}
	if err = result.Decode(&user); err != nil {
		return user, err
	}
	return user, err
}
