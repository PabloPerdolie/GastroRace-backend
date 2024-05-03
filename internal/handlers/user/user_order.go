package userhand

import (
	"backend/internal/handlers/middleware"
	cont "github.com/gorilla/context"
	"log"
	"net/http"
)

func AddOrder(w http.ResponseWriter, r *http.Request) {
	ctx := cont.Get(r, "user")
	user, ok := ctx.(middleware.UserData)

	log.Println(user)

	if !ok {
		http.Error(w, "Failed to transform", http.StatusBadRequest)
		return
	}
	// todo insert
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {

}
