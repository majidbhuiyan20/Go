package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"GO/model"
)

var students []model.Student

// GET all students
func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// POST a new student
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student model.Student

	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the student
	student.ID = strconv.FormatInt(time.Now().UnixNano(), 10)

	// Add to slice
	students = append(students, student)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}
