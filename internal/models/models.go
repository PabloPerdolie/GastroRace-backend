package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"desc" bson:"desc"`
	Price       string             `json:"price" bson:"price"`
	Type        string             `json:"type" bson:"type"`
	ImageId     interface{}        `json:"image_id" bson:"image_id"`
	ImageData   []byte             `json:"image_data" bson:"-"`
}
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
	IsAdmin  bool               `json:"is_admin" bson:"is_admin"`
	Cart     Cart               `json:"cart" bson:"cart"`
}

type Cart struct {
	Products []primitive.ObjectID `json:"products" bson:"products"`
}

type OrdersList struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Orders []Order
}

type Order struct {
	Products  []OrderProduct `json:"products" bson:"products"`
	OrderDate time.Time      `json:"order_date" bson:"order_date"`
}

type OrderProduct struct {
	Name  string `json:"name" bson:"name"`
	Price int    `json:"price" bson:"price"`
}
