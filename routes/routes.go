package routes

import (
	"GO/handlers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router) {
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
}
