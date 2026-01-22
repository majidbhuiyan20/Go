package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"GO/model"
)

var products []model.Product

// Initialize some products
func init() {
	products = append(products, model.Product{
		ID:          "1",
		Name:        "Iphone 16 pro max",
		Description: "This is a sample product",
		Price:       999.99,
		Image:       "uploads/iphone.jpg", // place sample image if needed
	})
}

// GET all products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// POST product (handles both multipart/form-data with file OR JSON with image path)
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "multipart/form-data") {
		handleMultipartProduct(w, r)
	} else {
		handleJSONProduct(w, r)
	}
}

// handle multipart/form-data upload
func handleMultipartProduct(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Invalid multipart/form-data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	priceStr := r.FormValue("price")

	if name == "" || priceStr == "" {
		http.Error(w, "Name and price are required", http.StatusBadRequest)
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}

	// Get image file
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Image required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Ensure uploads folder exists
	os.MkdirAll("uploads", os.ModePerm)

	// Unique image name
	imageName := strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(handler.Filename)
	imagePath := "uploads/" + imageName

	dst, err := os.Create(imagePath)
	if err != nil {
		http.Error(w, "Cannot save image", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Cannot save image", http.StatusInternalServerError)
		return
	}

	product := model.Product{
		ID:          strconv.FormatInt(time.Now().UnixNano(), 10),
		Name:        name,
		Description: description,
		Price:       price,
		Image:       imagePath,
	}

	products = append(products, product)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// handle JSON input
func handleJSONProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if product.Name == "" || product.Price == 0 {
		http.Error(w, "Name and price are required", http.StatusBadRequest)
		return
	}

	// Generate unique ID
	product.ID = strconv.FormatInt(time.Now().UnixNano(), 10)

	products = append(products, product)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
