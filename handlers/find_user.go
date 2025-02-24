package handlers

import (
	"api-simple-marketplace/db"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type UserExistsHandler struct{}

type UserExistsResult struct {
	UserExists bool
}

func (h *UserExistsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Should Use Get Method", http.StatusMethodNotAllowed)
		return
	}

	queryParam := r.URL.Query().Get("username")
	if queryParam == "" {
		fmt.Fprintf(w, "Query parameter 'username' is missing!")
		return
	}

	user := db.User{
		Username: queryParam,
	}

	userExists := true

	dbClient := r.Context().Value("db").(*gorm.DB)
	result := dbClient.First(&user)

	if result.Error != nil {
		userExists = false
	}

	userExistsResult := UserExistsResult{
		UserExists: userExists,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userExistsResult)
}
