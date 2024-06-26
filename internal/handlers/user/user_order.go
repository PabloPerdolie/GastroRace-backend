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
	"strconv"
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

	sum := 0
	for _, prod := range cart {
		i, _ := strconv.Atoi(prod.Price)
		sum += i
	}

	order := models.Order{
		ID:        primitive.NewObjectID(),
		UserId:    user.Id,
		Products:  cart,
		OrderDate: time.Now(),
		Status:    "CREATED",
		Sum:       sum,
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

func GetAllUserOrders(w http.ResponseWriter, r *http.Request) {
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

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	ctx := cont.Get(r, "user")
	user, ok := ctx.(middleware.UserData)

	if !ok {
		http.Error(w, "Failed to transform user data", http.StatusBadRequest)
		return
	}

	if !user.IsAdmin {
		http.Error(w, "Not enough rights", http.StatusLocked)
		return
	}

	orders, err := userrepo.GetAllOrders(context.Background())
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

func SetOrderStatus(w http.ResponseWriter, r *http.Request) {

}
