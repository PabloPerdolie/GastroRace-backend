package userrepo

import (
	"backend/internal/models"
	"backend/internal/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func AddNewOrder(ctx context.Context, order models.Order) error {
	order.ID = primitive.NewObjectID()
	one, err := mongodb.OrderColl.InsertOne(ctx, order)
	if err != nil {
		return err
	}
	if err = ClearCart(ctx, order.UserId); err != nil {
		return err
	}
	log.Printf("Inserted with id = %v", one.InsertedID)
	return nil
}

func GetOrders(ctx context.Context, userId primitive.ObjectID) ([]models.Order, error) {
	filter := bson.M{"user_id": userId}
	cursor, err := mongodb.OrderColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	orders := make([]models.Order, 1)
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func GetAllOrders(ctx context.Context) ([]models.Order, error) {
	cursor, err := mongodb.OrderColl.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

//todo status change
