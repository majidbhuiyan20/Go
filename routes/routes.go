package routes

import (
	"GO/handlers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router) {

	// ---------------- Category Routes ----------------
	r.HandleFunc("/admin/categories", handlers.CreateCategory).Methods("POST")
	r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")
	r.HandleFunc("/createstudent", handlers.CreateStudentList).Methods("POST")
	r.HandleFunc("/students", handlers.GetStudent).Methods("GET")

	r.HandleFunc("/admin/products", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/by-category", handlers.GetProductsByCategory).Methods("GET")

}
