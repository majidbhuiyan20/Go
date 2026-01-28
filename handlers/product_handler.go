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

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var prod model.Product

	err := json.NewDecoder(r.Body).Decode(&prod)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	if prod.Name == "" || prod.CategoryID == "" || prod.Price == 0 || prod.Stock < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			Success: false,
			Message: "Invalid Product data, name, category_id, price(>0)",
		})
		return

	}
	var id int

	err = config.DB.QueryRow(`Insert into products(category_id, name, description, price, stock, image, created_at)
		values ($1, $2, $3, $4, $5, $6, $7) returning id`,
		prod.CategoryID, prod.Name, prod.Description, prod.Price, prod.Stock, prod.Image, time.Now().Unix(),
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
	prod.ID = strconv.Itoa(id)
	prod.CreatedAt = time.Now().Unix()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.Response{
		Success: true,
		Message: "Product Created Successfully",
	})
}

func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categoryIDStr := r.URL.Query().Get("category_id")
	if categoryIDStr == "" {
		http.Error(w, "category_id is required", http.StatusBadRequest)
		return
	}

	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		http.Error(w, "invalid category_id", http.StatusBadRequest)
		return
	}

	rows, err := config.DB.Query(`
        SELECT id, category_id, name, description, price, stock, image, created_at
        FROM products
        WHERE category_id = $1
    `, categoryID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	products := []model.Product{}

	for rows.Next() {
		var p model.Product
		rows.Scan(
			&p.ID,
			&p.CategoryID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Stock,
			&p.Image,
			&p.CreatedAt,
		)
		products = append(products, p)
	}

	json.NewEncoder(w).Encode(products)
}
