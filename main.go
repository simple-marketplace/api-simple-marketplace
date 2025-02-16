package main

import (
	"api-simple-marketplace/db"
	"api-simple-marketplace/es"
	"api-simple-marketplace/handlers"
	"api-simple-marketplace/middleware"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	dsn := "host=icy-smoke-2219.fly.dev user=postgres password=df05vLaFDnmefci dbname=postgres port=5433 sslmode=disable"

	dbMiddleware := middleware.DBMiddleware{
		M: db.NewDatabaseClient(dsn),
	}

	esMiddleware := middleware.ESMiddleware{
		M: es.NewElasticsearchClient(),
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// _db.AutoMigrate(&db.Product{})

	mux := http.NewServeMux()
	mux.HandleFunc("/products", dbMiddleware.Apply(handlers.GetProductsHandler{}))
	mux.HandleFunc("/products/create", dbMiddleware.Apply(handlers.CreateProductHandler{}))
	mux.HandleFunc("/products/search", esMiddleware.Apply(handlers.SearchProductsHandler{}))

	muxWithCors := c.Handler(mux)
	http.ListenAndServe(":8080", muxWithCors)

}
