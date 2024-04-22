package handlers

import (
	producthand "backend/internal/handlers/product"
	userhand "backend/internal/handlers/user"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://gastrorace-frontend.onrender.com"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
	})

	r := mux.NewRouter()
	r.Use(c.Handler)

	http.Handle("/", r)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", userhand.SignUp).Methods("POST")
	auth.HandleFunc("/signin", userhand.SignIn).Methods("POST")

	api := r.PathPrefix("/api/v1").Subrouter()
	//api.Use(middleware.Authorize)

	cart := api.PathPrefix("/cart").Subrouter()
	cart.HandleFunc("", userhand.GetCart).Methods("GET")
	cart.HandleFunc("/remove", userhand.RemoveFromCart).Methods("POST")
	cart.HandleFunc("/add", userhand.AddToCart).Methods("POST")

	products := api.PathPrefix("/products").Subrouter()
	products.HandleFunc("", producthand.Create).Methods("POST")
	products.HandleFunc("", producthand.GetAll).Methods("GET")
	// todo DeleteProduct

	// todo AddOrder
	// todo GetOrders

}
