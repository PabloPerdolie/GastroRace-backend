package userhand

import (
	"backend/internal/models"
	"backend/internal/security"
	"encoding/json"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Unable to decode data", http.StatusBadRequest)
		return
	}
	_, err := security.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusBadRequest)
		return
	}
	token, err := security.GenerateToken(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Unable to decode data", http.StatusBadRequest)
		return
	}

	token, err := security.GenerateToken(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}
