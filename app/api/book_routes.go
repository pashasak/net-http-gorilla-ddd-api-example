package api

import (
	"context"
	"github.com/gorilla/mux"
	"net-http-gorilla-ddd-example/app/middleware"
)

func NewBookRoutes(ctx context.Context, router *mux.Router, app *Book) *mux.Router {
	// API Version 1 Router
	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	middleware.RegisterMiddleware(ctx, apiV1)
	apiV1.HandleFunc("/books", app.GetBooks).Methods("GET")
	apiV1.HandleFunc("/book/{id}", app.GetBook).Methods("GET")
	apiV1.HandleFunc("/book/{id}", app.UpdateBook).Methods("PUT")
	apiV1.HandleFunc("/book", app.CreateBook).Methods("POST")
	apiV1.HandleFunc("/book/{id}", app.DeleteBook).Methods("DELETE")
	return apiV1
}
