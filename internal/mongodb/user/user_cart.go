package userrepo

import (
	"backend/internal/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUserCart(ctx context.Context, userId primitive.ObjectID, productId primitive.ObjectID) error {
	user, err := FindByIdUser(context.Background(), userId)
	if err != nil {
		return err
	}
	user.Cart.Products = append(user.Cart.Products, productId)

	_, err = mongodb.UserColl.UpdateByID(context.Background(), userId, user)
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
			user.Cart.Products = append(user.Cart.Products[:i], user.Cart.Products[i:]...)
			break
		}
	}

	_, err = mongodb.UserColl.UpdateByID(context.Background(), userId, user)
	if err != nil {
		return err
	}
	return nil
}
