package middleware

import (
	"context"
	"net/http"

	"gorm.io/gorm"
)

type DBMiddleware struct {
	M *gorm.DB
}

func (m *DBMiddleware) Apply(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new context with the DB client
		ctx := context.WithValue(r.Context(), "db", m.M)

		// Pass the context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
