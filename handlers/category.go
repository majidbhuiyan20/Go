package handlers

import (
	"GO/config"
	"GO/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

// ---------------- Admin ----------------
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cat model.Category

	err := json.NewDecoder(r.Body).Decode(&cat)
	log.Println("Received Category:", cat)

	if err != nil || cat.Name == "" {
		http.Error(w, "Invalid Category Data", http.StatusBadRequest)
		return
	}

	// Insert into PostgreSQL
	var id int
	err = config.DB.QueryRow(
		`INSERT INTO categories (name, image, created_at)
		 VALUES ($1, $2, $3) RETURNING id`,
		cat.Name, cat.Image, time.Now().Unix(),
	).Scan(&id)

	if err != nil {
		http.Error(w, "Failed to insert category", http.StatusInternalServerError)
		log.Println("Insert error:", err)
		return
	}

	cat.ID = strconv.Itoa(id)
	cat.CreatedAt = time.Now().Unix()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cat)
}

// ---------------- End User ----------------
func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := config.DB.Query(`SELECT id, name, image, created_at FROM categories`)
	if err != nil {
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
		log.Println("Query error:", err)
		return
	}
	defer rows.Close()

	categories := []model.Category{}

	for rows.Next() {
		var cat model.Category
		var id int
		err := rows.Scan(&id, &cat.Name, &cat.Image, &cat.CreatedAt)
		if err != nil {
			http.Error(w, "Failed to scan category", http.StatusInternalServerError)
			log.Println("Scan error:", err)
			return
		}
		cat.ID = strconv.Itoa(id)
		categories = append(categories, cat)
	}

	json.NewEncoder(w).Encode(categories)
}
