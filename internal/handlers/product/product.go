package product

import (
	"backend/internal/models"
	productrepo "backend/internal/mongodb/product"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	formFile, _, err := r.FormFile("file") // "file" - имя поля формы для файла
	if err != nil {
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(formFile)
	if err != nil {
		http.Error(w, "Unable to decode file", http.StatusBadRequest)
		return
	}

	var prod models.Product
	if err := json.NewDecoder(r.Body).Decode(&prod); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	prod.ImageData = bytes

	if err := productrepo.CreateProduct(context.Background(), prod); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Data has been succecfully saved in MongoDB")))
}
