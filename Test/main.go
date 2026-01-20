package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Student model
type Student struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// In-memory student data
var students = []Student{
	{Name: "Majid Bhuiyan", ID: "101"},
}

// GET all students
func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// POST new student
func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newStudent Student
	err := json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	students = append(students, newStudent)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newStudent)
}

func main() {
	r := mux.NewRouter()

	// Route GET request
	r.HandleFunc("/students", getStudents).Methods("GET")

	// Route POST request
	r.HandleFunc("/students", createStudent).Methods("POST")

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
