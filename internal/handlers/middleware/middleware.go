package middleware

import (
	"backend/internal/security"
	cont "github.com/gorilla/context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		if header == "" {
			http.Error(w, "Empty authorization header", http.StatusUnauthorized)
			return
		}
		headerPart := strings.Split(header, " ")
		if len(headerPart) != 2 {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}

		userId, err := security.ParseToken(headerPart[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		user, err := security.GetUser(userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		cont.Set(r, "user", UserData{
			Id:      user.ID,
			IsAdmin: user.IsAdmin,
		})
		next.ServeHTTP(w, r)
	})
}

type UserData struct {
	Id      primitive.ObjectID
	IsAdmin bool
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		//w.Header().Set("Access-Control-Allow-Origin", "https://gastrorace-frontend.onrender.com")
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
