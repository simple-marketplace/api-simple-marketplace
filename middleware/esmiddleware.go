package middleware

import (
	"context"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
)

type ESMiddleware struct {
	M *elasticsearch.Client
}

func (m *ESMiddleware) Apply(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "es", m.M)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
