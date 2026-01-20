package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Student struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

var students = []Student{
	{Name: "Majid Bhuiyan", ID: "101"},
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(students)

	case http.MethodPost:
		var newStudent Student
		err := json.NewDecoder(r.Body).Decode(&newStudent)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		students = append(students, newStudent)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newStudent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/students", studentsHandler)
	log.Println("Server running os http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
