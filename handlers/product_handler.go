package handlers

import (
	"GO/config"
	"GO/model"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var prod model.Product

	err := json.NewDecoder(r.Body).Decode(&prod)

	if err != nil || prod.Name == "" || prod.CategoryID == "" || prod.Price <= 0 || prod.Stock < 0 {
		http.Error(w, "Invalid Product Data", http.StatusBadRequest)
		return
	}

	var id int

	err = config.DB.QueryRow(
		`INSERT INTO products (category_id, name, description, price, stock, image, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		prod.CategoryID, prod.Name, prod.Description, prod.Price, prod.Stock, prod.Image, time.Now().Unix(),
	).Scan(&id)

	if err != nil {
		http.Error(w, "Failed to insert product", http.StatusInternalServerError)
		return
	}

	prod.ID = strconv.Itoa(id)
	prod.CreatedAt = time.Now().Unix()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(prod)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := config.DB.Query(`SELECT id, category_id, name, description, price, stock, image, created_at FROM products`)
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	products := []model.Product{}

	for rows.Next() {
		var prod model.Product
		err := rows.Scan(&prod.ID, &prod.CategoryID, &prod.Name, &prod.Description, &prod.Price, &prod.Stock, &prod.Image, &prod.CreatedAt)
		if err != nil {
			http.Error(w, "Error scanning product", http.StatusInternalServerError)
			return
		}
		products = append(products, prod)
	}

	json.NewEncoder(w).Encode(products)
}
