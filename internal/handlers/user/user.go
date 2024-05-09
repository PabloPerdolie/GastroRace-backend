package userhand

import (
	"backend/internal/models"
	"backend/internal/security"
	"encoding/json"
	"log"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
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
	token, err, _ := security.GenerateToken(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Unable to decode data", http.StatusBadRequest)
		return
	}

	token, err, isAdmin := security.GenerateToken(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(struct {
		Token   string `json:"token"`
		IsAdmin bool   `json:"IsAdmin"`
	}{
		Token:   token,
		IsAdmin: isAdmin,
	})
	if err != nil {
		http.Error(w, "Failed to marshal", http.StatusInternalServerError)
		return
	}
	log.Println()

	log.Println("Successful sign in with token=" + token)

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
