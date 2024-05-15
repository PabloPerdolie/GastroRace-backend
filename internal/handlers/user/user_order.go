package userhand

import (
	"backend/internal/handlers/middleware"
	"backend/internal/models"
	userrepo "backend/internal/mongodb/user"
	"context"
	"encoding/json"
	cont "github.com/gorilla/context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func AddOrder(w http.ResponseWriter, r *http.Request) {
	ctx := cont.Get(r, "user")
	user, ok := ctx.(middleware.UserData)

	if !ok {
		http.Error(w, "Failed to transform user data", http.StatusBadRequest)
		return
	}

	cart, err := userrepo.GetCart(context.Background(), user.Id)
	if err != nil {
		http.Error(w, "Failed to get cart", http.StatusBadRequest)
		return
	}

	if len(cart) == 0 {
		http.Error(w, "Cart is empty", http.StatusBadRequest)
		return
	}

	order := models.Order{
		ID:        primitive.NewObjectID(),
		UserId:    user.Id,
		Products:  cart,
		OrderDate: time.Now(),
		Status:    "CREATED",
	}

	err = userrepo.AddNewOrder(context.Background(), order)
	if err != nil {
		http.Error(w, "Failed to add new order to MongoDB", http.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(order)
	if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	ctx := cont.Get(r, "user")
	user, ok := ctx.(middleware.UserData)

	if !ok {
		http.Error(w, "Failed to transform user data", http.StatusBadRequest)
		return
	}

	orders, err := userrepo.GetOrders(context.Background(), user.Id)
	if err != nil {
		http.Error(w, "Failed to get orders from MongoDB", http.StatusBadRequest)
		return
	}
	bytes, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

//todo status change
