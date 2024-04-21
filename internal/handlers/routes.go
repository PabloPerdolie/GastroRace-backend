package handlers

import (
	"backend/internal/handlers/product"
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

	r.HandleFunc("/product", product.Create).Methods("POST")
	r.HandleFunc("/products", product.GetAll).Methods("GET")
}
