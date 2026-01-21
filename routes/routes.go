package routes

import (
	"GO/handlers"
	"github.com/gorilla/mux"
)

// RegisterProductRoutes registers product and student routes
func RegisterProductRoutes(r *mux.Router) {
	// Product routes
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")

	// Student routes
	r.HandleFunc("/students", handlers.GetStudents).Methods("GET")
	r.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
}
