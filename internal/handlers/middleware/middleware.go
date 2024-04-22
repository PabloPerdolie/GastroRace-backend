package middleware

import (
	"backend/internal/security"
	"context"
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

		ctx := context.WithValue(context.Background(), "isAdmin", user.IsAdmin)
		r.WithContext(ctx)
		//context.Set(r, USERROLE, user.Role)
		next.ServeHTTP(w, r)
	})
}
