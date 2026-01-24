package handlers

import (
	"GO/model"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

var categories []model.Category

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cat model.Category

	err := json.NewDecoder(r.Body).Decode(&cat)

	if err != nil || cat.Name == "" {
		http.Error(w, "Invalid Category Data", http.StatusBadRequest)
		return
	}
	cat.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	cat.CreatedAt = time.Now().Unix()

	categories = append(categories, cat)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cat)
}
