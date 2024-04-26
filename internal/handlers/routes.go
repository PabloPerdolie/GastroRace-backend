package handlers

import (
	"backend/internal/handlers/middleware"
	producthand "backend/internal/handlers/product"
	userhand "backend/internal/handlers/user"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "https://gastrorace-frontend.onrender.com")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}
			// Продолжить выполнение цепочки обработчиков
			next.ServeHTTP(w, r)
		})
	}
	r.Use(corsMiddleware)

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/signup", userhand.SignUp).Methods("POST", "OPTIONS")
	auth.HandleFunc("/signin", userhand.SignIn).Methods("POST", "OPTIONS")

	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.Authorize)

	cart := api.PathPrefix("/cart").Subrouter()
	cart.HandleFunc("", userhand.GetCart).Methods("GET")
	cart.HandleFunc("/remove", userhand.RemoveFromCart).Methods("POST")
	cart.HandleFunc("/add", userhand.AddToCart).Methods("POST")

	products := api.PathPrefix("/products").Subrouter()
	products.HandleFunc("", producthand.Create).Methods("POST", "OPTIONS")
	products.HandleFunc("", producthand.GetAll).Methods("GET", "OPTIONS")
	// todo DeleteProduct

	// todo AddOrder
	// todo GetOrders

	return r
}
