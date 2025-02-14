package middleware

import (
	"context"
	"net/http"

	"gorm.io/gorm"
)

func DBMiddleware(_db *gorm.DB, next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new context with the DB client
		ctx := context.WithValue(r.Context(), "db", _db)

		// Pass the context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
