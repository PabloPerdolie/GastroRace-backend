package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CartProductDTO struct {
	Id    primitive.ObjectID `json:"id"`
	Name  string             `json:"name"`
	Price string             `json:"price"`
}
