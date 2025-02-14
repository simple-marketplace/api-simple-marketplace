package main

import (
	"api-simple-marketplace/db"
	"api-simple-marketplace/es"
	"api-simple-marketplace/handlers"
	"api-simple-marketplace/middleware"
	"net/http"
)

func main() {
	dsn := "host=icy-smoke-2219.fly.dev user=postgres password=df05vLaFDnmefci dbname=postgres port=5433 sslmode=disable"
	_db := db.NewDatabaseClient(dsn)
	_es := es.NewElasticsearchClient()

	// _db.AutoMigrate(&db.Product{})

	mux := http.NewServeMux()
	mux.HandleFunc("/products", middleware.DBMiddleware(_db, http.HandlerFunc(handlers.GetProducts)))
	mux.HandleFunc("/products/create", middleware.DBMiddleware(_db, http.HandlerFunc(handlers.CreateProduct)))
	mux.HandleFunc("/products/search", middleware.ESMiddleware(_es, http.HandlerFunc(handlers.SearchProducts)))
	http.ListenAndServe(":8080", mux)

}
