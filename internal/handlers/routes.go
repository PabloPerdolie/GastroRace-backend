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
		AllowedOrigins: []string{"http://localhost:3000", "https://gastrorace-frontend.onrender.com"},
	})

	r := mux.NewRouter()
	r.Use(c.Handler)

	http.Handle("/", r)

	auth := r.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/signup", userhand.SignUp).Methods("POST")
	auth.HandleFunc("/signin", userhand.SignIn).Methods("POST")

	r.HandleFunc("/product", producthand.Create).Methods("POST")
	r.HandleFunc("/products", producthand.GetAll).Methods("GET")

}
