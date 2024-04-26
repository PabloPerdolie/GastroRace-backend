package producthand

import (
	"backend/internal/models"
	productrepo "backend/internal/mongodb/product"
	bytes2 "bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	formFile, _, err := r.FormFile("file") // "file" - имя поля формы для файла
	if err != nil {
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}

	//bytes, err := json.Marshal(formFile)
	buf := bytes2.NewBuffer(nil)
	if _, err := io.Copy(buf, formFile); err != nil {
		http.Error(w, "Unable to retrieve file", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "Unable to decode file", http.StatusBadRequest)
		return
	}

	prod := models.Product{
		Name:        r.FormValue("name"),
		Description: r.FormValue("desc"),
		Price:       r.FormValue("price"),
		Type:        r.FormValue("type"),
		ImageData:   buf.Bytes(),
	}

	if err := productrepo.CreateProduct(context.Background(), prod); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Data has been successfully saved in MongoDB")))
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	all, err := productrepo.FindAll(context.Background())
	if err != nil {
		http.Error(w, "Failed to find products", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	bytes, err := json.Marshal(all)
	if err != nil {
		http.Error(w, "Unable to decode file", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
