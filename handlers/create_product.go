package handlers

import (
	"api-simple-marketplace/db"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Should Use Post Method", http.StatusMethodNotAllowed)
		return
	}

	dbClient := r.Context().Value("db").(*gorm.DB)

	product := &db.Product{}
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res := dbClient.Create(&product)

	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("%v\n", res)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.Result{
		ID: product.ID,
	})
}
