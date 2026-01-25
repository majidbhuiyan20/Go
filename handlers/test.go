package handlers

import (
	"GO/config"
	"GO/model"
	"encoding/json"
	"net/http"
	"strconv"
)

// ==============This is Admin Section==================//
func CreateStudentList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var majid model.StudentInfo

	// Decode JSON body
	err := json.NewDecoder(r.Body).Decode(&majid)
	if err != nil || majid.Name == "" || majid.ClassName == "" || majid.RollNo == "" {
		http.Error(w, "Invalid Student Data", http.StatusBadRequest)
		return
	}

	// Insert into database
	var id int64
	err = config.DB.QueryRow(
		`INSERT INTO students (name, school, class_name, roll_no)
     VALUES ($1, $2, $3, $4)
     RETURNING id`,
		majid.Name, majid.School, majid.ClassName, majid.RollNo,
	).Scan(&id)

	if err != nil {
		http.Error(w, "Failed to insert student: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare response
	majid.ID = strconv.FormatInt(id, 10)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(majid)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := config.DB.Query("SELECT id, name, school, class_name, roll_no FROM students")
	if err != nil {
		http.Error(w, "Failed to fetch students: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []model.StudentInfo

	for rows.Next() {
		var s model.StudentInfo
		err := rows.Scan(&s.ID, &s.Name, &s.School, &s.ClassName, &s.RollNo)
		if err != nil {
			http.Error(w, "Failed to read student data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		students = append(students, s)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}
