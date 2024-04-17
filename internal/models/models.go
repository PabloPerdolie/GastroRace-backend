package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"desc" bson:"desc"`
	Price       int                `json:"price" bson:"price"`
	Type        string             `json:"type" bson:"type"`
	ImageId     interface{}        `json:"-" bson:"image_id"`
	ImageData   []byte             `json:"-" bson:"-"`
}
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
	IsAdmin  bool               `json:"is_admin" bson:"is_admin"`
}

type Cart struct {
}

type OrdersList struct {
}