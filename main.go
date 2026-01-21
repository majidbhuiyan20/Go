package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	ID     string `json:"ID"`
	Name   string `json:"Name"`
	School string `json:"School"`
}

var students = []Student{
	{Name: "Majid Bhuiyan", ID: "101", School: "Jawar High School"},
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(students)
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

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

	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/students", createStudent).Methods("POST")
	log.Println("Server Running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
