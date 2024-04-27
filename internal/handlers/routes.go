package handlers

import (
	"backend/internal/handlers/middleware"
	producthand "backend/internal/handlers/product"
	userhand "backend/internal/handlers/user"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.Cors)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", userhand.SignUp).Methods("POST", "OPTIONS")
	auth.HandleFunc("/signin", userhand.SignIn).Methods("POST", "OPTIONS")

	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.Authorize)

	cart := api.PathPrefix("/cart").Subrouter()
	cart.HandleFunc("", userhand.GetCart).Methods("GET", "OPTIONS")
	cart.HandleFunc("/remove", userhand.RemoveFromCart).Methods("POST", "OPTIONS")
	cart.HandleFunc("/add", userhand.AddToCart).Methods("POST", "OPTIONS")

	products := api.PathPrefix("/products").Subrouter()
	products.HandleFunc("", producthand.Create).Methods("POST", "OPTIONS")
	products.HandleFunc("", producthand.GetAll).Methods("GET", "OPTIONS")
	// todo DeleteProduct

	// todo AddOrder
	// todo GetOrders

	return r
}
