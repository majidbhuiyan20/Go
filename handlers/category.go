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

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cat model.Category

	err := json.NewDecoder(r.Body).Decode(&cat)

	if err != nil {
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Invalid Reques Body",
		})
		return
	}

	if cat.Name == "" {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Category name is required",
		})
		return
	}

	var id int
	err = config.DB.QueryRow(
		`INSERT INTO categories (name, image, created_at)
		 VALUES ($1, $2, $3) RETURNING id`,
		cat.Name, cat.Image, time.Now().Unix(),
	).Scan(&id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Insert error: ", err)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Failed to create category",
		})
		return
	}

	cat.ID = strconv.Itoa(id)
	cat.CreatedAt = time.Now().Unix()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.Response{
		Success: true,
		Message: "Category created Successfully",
	})

}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := config.DB.Query(`SELECT id, name, image, created_at from categories order by id`)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Query error:", err)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Failed to fetch catrgories",
		})
		return
	}
	defer rows.Close()

	categories := []model.Category{}

	for rows.Next() {
		var cat model.Category
		var id int
		err := rows.Scan(&id, &cat.Name, &cat.Image, &cat.CreatedAt)

		if err != nil {
			w.WriteHeader((http.StatusInternalServerError))
			log.Println("Scan error: ", err)
			json.NewEncoder(w).Encode(model.Response{
				Success: false,
				Message: "Failed to scan category",
			})
			return
		}
		cat.ID = strconv.Itoa(id)
		categories = append(categories, cat)
	}
	json.NewEncoder(w).Encode(model.Response{
		Success: true,
		Message: "Categories fetch successfully",
		Data:    categories,
	})
}
