package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"GO/model"
)

var students []model.Student

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student model.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	student.ID = strconv.FormatInt(time.Now().UnixNano(), 10)

	students = append(students, student)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}
