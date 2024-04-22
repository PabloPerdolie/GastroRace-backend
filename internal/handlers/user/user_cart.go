package userhand

import (
	"fmt"
	"net/http"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// todo extract userId from token to update cart by id
	//userrepo.UpdateUserCart(context.Background(),)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Data has been successfully saved in MongoDB")))
}

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Data has been successfully saved in MongoDB")))
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("")))
}
