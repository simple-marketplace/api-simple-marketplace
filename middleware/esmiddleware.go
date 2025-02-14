package middleware

import (
	"context"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
)

func ESMiddleware(_es *elasticsearch.Client, next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "es", _es)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
