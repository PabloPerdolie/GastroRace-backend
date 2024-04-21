package userhand

import (
	"backend/internal/models"
	userrepo "backend/internal/mongodb/user"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Unable to decode data", http.StatusBadRequest)
		return
	}
	err := userrepo.CreateUser(context.Background(), user)
	if err != nil {
		http.Error(w, "Failed to create document in MongoDB", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Data has been successfully saved in MongoDB")))
}

func SignIn(w http.ResponseWriter, r *http.Request) {

}
