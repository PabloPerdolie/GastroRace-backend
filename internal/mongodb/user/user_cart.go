package userrepo

import (
	"backend/internal/models"
	"backend/internal/mongodb"
	productrepo "backend/internal/mongodb/product"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCart(ctx context.Context, userId primitive.ObjectID) ([]models.Product, error) {
	user, err := FindByIdUser(ctx, userId)
	if err != nil {
		return []models.Product{}, err
	}
	cart := make([]models.Product, len(user.Cart.Products))
	for i, prodId := range user.Cart.Products {
		one, err := productrepo.FindOne(context.Background(), prodId)
		if err != nil {
			return nil, err
		}
		one.ImageData = nil
		cart[i] = one
	}
	return cart, err
}

func UpdateUserCart(ctx context.Context, userId primitive.ObjectID, productId primitive.ObjectID) error {
	user, err := FindByIdUser(context.Background(), userId)
	if err != nil {
		return err
	}
	user.Cart.Products = append(user.Cart.Products, productId)

	update := bson.M{
		"$set": bson.M{
			"cart.products": user.Cart.Products,
		},
	}

	_, err = mongodb.UserColl.UpdateByID(ctx, userId, update)
	if err != nil {
		return err
	}
	return nil
}

func RemoveFromUserCart(ctx context.Context, userId primitive.ObjectID, productId primitive.ObjectID) error {
	user, err := FindByIdUser(context.Background(), userId)
	if err != nil {
		return err
	}
	for i, prod := range user.Cart.Products {
		if prod == productId {
			user.Cart.Products = append(user.Cart.Products[:i], user.Cart.Products[i+1:]...)
			break
		}
	}
	update := bson.M{
		"$set": bson.M{
			"cart.products": user.Cart.Products,
		},
	}
	_, err = mongodb.UserColl.UpdateByID(context.Background(), userId, update)
	if err != nil {
		return err
	}
	return nil
}
