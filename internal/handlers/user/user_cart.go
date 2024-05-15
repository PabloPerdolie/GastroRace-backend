package userhand

import (
	"backend/internal/handlers/middleware"
	userrepo "backend/internal/mongodb/user"
	"context"
	"encoding/json"
	cont "github.com/gorilla/context"
	"log"

	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := cont.Get(r, "user")
	user, ok := ctx.(middleware.UserData)

	log.Println(user)

	if !ok {
		http.Error(w, "Failed to transform", http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Unable to transform id form string to ObjectId", http.StatusBadRequest)
		return
	}

	err = userrepo.UpdateUserCart(context.Background(), user.Id, hex)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to save data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Data has been successfully saved in MongoDB")))
}

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := cont.Get(r, "user")
	user, ok := ctx.(middleware.UserData)

	log.Println(user)

	if !ok {
		http.Error(w, "Failed to transform", http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")
	hex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Unable to transform id form string to ObjectId", http.StatusBadRequest)
		return
	}

	err = userrepo.RemoveFromUserCart(context.Background(), user.Id, hex)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to save data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Data has been successfully removed from MongoDB")))
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := cont.Get(r, "user")
	user, ok := ctx.(middleware.UserData)

	log.Println(user)

	if !ok {
		http.Error(w, "Failed to transform", http.StatusBadRequest)
		return
	}
	cart, err := userrepo.GetCart(context.Background(), user.Id)
	if err != nil {
		http.Error(w, "Unable to decode file", http.StatusInternalServerError)
		return
	}
	log.Println(cart)
	bytes, err := json.Marshal(cart)
	if err != nil {
		http.Error(w, "Unable to decode file", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
