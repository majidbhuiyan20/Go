package routes

import (
	"GO/handlers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router) {

	// ---------------- Category Routes ----------------
	r.HandleFunc("/admin/categories", handlers.CreateCategory).Methods("POST")

	r.HandleFunc("/categories", handlers.GetCategories).Methods("GET")

	// ---------------- Product Routes ----------------
	//r.HandleFunc("/admin/categories/{categoryId}/products", handlers.CreateProduct).Methods("POST")

}

// ------Api End POint Routes------
/*
Category APIs

Admin Panel

POST http://localhost:8080/admin/categories → Create a new category

PUT http://localhost:8080/admin/categories/{id} → Update category name or image

DELETE http://localhost:8080/admin/categories/{id} → Delete a category

End User
4. GET http://localhost:8080/categories → Get all categories
5. GET http://localhost:8080/categories/{id} → Get category by ID (including products)

Product APIs (Category-wise)

Admin Panel
6. POST http://localhost:8080/admin/categories/{categoryId}/products → Create a product under category (JSON or multipart/form-data)
7. PUT http://localhost:8080/admin/categories/{categoryId}/products/{productId} → Update product details
8. DELETE http://localhost:8080/admin/categories/{categoryId}/products/{productId} → Delete a product

End User
9. GET http://localhost:8080/categories/{categoryId}/products → Get all products under category

Optional query: ?active=true → only active products
*/ //
