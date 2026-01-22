package main

import (
	"log"
	"net/http"
	"os"

	"GO/routes"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()

	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	routes.RegisterProductRoutes(r)

	log.Printf("Server running on port http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
