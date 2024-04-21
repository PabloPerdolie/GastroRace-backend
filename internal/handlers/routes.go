package handlers

import (
	"backend/internal/handlers/product"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes() {
	r := mux.NewRouter()
	http.Handle("/", r)

	r.HandleFunc("/product", product.Create).Methods("POST")
	r.HandleFunc("/products", product.GetAll).Methods("GET")
}
