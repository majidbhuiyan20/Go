package main

import (
	"log"
	"net/http"

	"GO/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Serve images
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	routes.RegisterProductRoutes(r)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
