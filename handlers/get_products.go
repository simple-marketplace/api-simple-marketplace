package handlers

import (
	"api-simple-marketplace/db"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type GetProductsHandler struct{}

func (h GetProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Should Use Get Method", http.StatusMethodNotAllowed)
		return
	}

	var product db.Product

	dbClient := r.Context().Value("db").(*gorm.DB)
	dbClient.Find(&product, "name = ?", "eggs")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
