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

	"github.com/gorilla/mux"
)

// ---------------- Admin ----------------

// Create product under category
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categoryId := mux.Vars(r)["categoryId"]

	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "multipart/form-data") {
		handleMultipartProductByCategory(w, r, categoryId)
	} else {
		handleJSONProductByCategory(w, r, categoryId)
	}
}

// Update Product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categoryId := mux.Vars(r)["categoryId"]
	productId := mux.Vars(r)["productId"]

	var updateProd model.Product
	contentType := r.Header.Get("Content-Type")

	if strings.HasPrefix(contentType, "multipart/form-data") {
		r.ParseMultipartForm(20 << 20)
		updateProd.Name = r.FormValue("name")
		updateProd.Description = r.FormValue("description")
		if r.FormValue("price") != "" {
			updateProd.Price, _ = strconv.ParseFloat(r.FormValue("price"), 64)
		}
		if r.FormValue("stock") != "" {
			updateProd.Stock, _ = strconv.Atoi(r.FormValue("stock"))
		}
		updateProd.SKU = r.FormValue("sku")
		updateProd.Tags = strings.Split(r.FormValue("tags"), ",")
		updateProd.IsFeatured = strings.ToLower(r.FormValue("is_featured")) == "true"
		updateProd.IsActive = strings.ToLower(r.FormValue("is_active")) == "true"

		file, handler, err := r.FormFile("image")
		if err == nil {
			defer file.Close()
			os.MkdirAll("uploads/products", os.ModePerm)
			imageName := strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(handler.Filename)
			imagePath := "uploads/products/" + imageName
			dst, _ := os.Create(imagePath)
			defer dst.Close()
			io.Copy(dst, file)
			updateProd.Image = imagePath
		} else {
			updateProd.Image = r.FormValue("image_url")
		}
	} else {
		json.NewDecoder(r.Body).Decode(&updateProd)
	}

	for i, cat := range categories {
		if cat.ID == categoryId {
			for j, prod := range cat.Products {
				if prod.ID == productId {
					if updateProd.Name != "" {
						categories[i].Products[j].Name = updateProd.Name
					}
					if updateProd.Description != "" {
						categories[i].Products[j].Description = updateProd.Description
					}
					if updateProd.Price != 0 {
						categories[i].Products[j].Price = updateProd.Price
					}
					if updateProd.Image != "" {
						categories[i].Products[j].Image = updateProd.Image
					}
					if updateProd.Stock != 0 {
						categories[i].Products[j].Stock = updateProd.Stock
					}
					if updateProd.SKU != "" {
						categories[i].Products[j].SKU = updateProd.SKU
					}
					if updateProd.Tags != nil {
						categories[i].Products[j].Tags = updateProd.Tags
					}
					categories[i].Products[j].IsFeatured = updateProd.IsFeatured
					categories[i].Products[j].IsActive = updateProd.IsActive
					categories[i].Products[j].UpdatedAt = time.Now().Unix()
					json.NewEncoder(w).Encode(categories[i].Products[j])
					return
				}
			}
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

// Delete product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	categoryId := mux.Vars(r)["categoryId"]
	productId := mux.Vars(r)["productId"]

	for i, cat := range categories {
		if cat.ID == categoryId {
			for j, prod := range cat.Products {
				if prod.ID == productId {
					categories[i].Products = append(categories[i].Products[:j], categories[i].Products[j+1:]...)
					w.WriteHeader(http.StatusNoContent)
					return
				}
			}
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
	}
	http.Error(w, "Category not found", http.StatusNotFound)
}

// ---------------- End User ----------------

// Get products by category
func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categoryId := mux.Vars(r)["categoryId"]

	for _, cat := range categories {
		if cat.ID == categoryId {
			activeOnly := r.URL.Query().Get("active")
			if activeOnly == "true" {
				filtered := []model.Product{}
				for _, p := range cat.Products {
					if p.IsActive {
						filtered = append(filtered, p)
					}
				}
				json.NewEncoder(w).Encode(filtered)
				return
			}
			json.NewEncoder(w).Encode(cat.Products)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

// ---------------- JSON Product ----------------
func handleJSONProductByCategory(w http.ResponseWriter, r *http.Request, categoryId string) {
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if product.Name == "" || product.Price == 0 {
		http.Error(w, "Name and price required", http.StatusBadRequest)
		return
	}

	product.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	product.CreatedAt = time.Now().Unix()
	product.UpdatedAt = time.Now().Unix()

	for i, cat := range categories {
		if cat.ID == categoryId {
			categories[i].Products = append(categories[i].Products, product)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

// ---------------- Multipart Product ----------------
func handleMultipartProductByCategory(w http.ResponseWriter, r *http.Request, categoryId string) {
	err := r.ParseMultipartForm(20 << 20)
	if err != nil {
		http.Error(w, "Invalid multipart/form-data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	priceStr := r.FormValue("price")
	stockStr := r.FormValue("stock")
	sku := r.FormValue("sku")
	tagsStr := r.FormValue("tags")
	isFeaturedStr := r.FormValue("is_featured")
	isActiveStr := r.FormValue("is_active")

	if name == "" || priceStr == "" {
		http.Error(w, "Name and price required", http.StatusBadRequest)
		return
	}

	price, _ := strconv.ParseFloat(priceStr, 64)
	stock := 0
	if stockStr != "" {
		stock, _ = strconv.Atoi(stockStr)
	}
	isFeatured := strings.ToLower(isFeaturedStr) == "true"
	isActive := strings.ToLower(isActiveStr) == "true"

	// handle image file or URL
	file, handler, err := r.FormFile("image")
	imagePath := r.FormValue("image_url")
	if err == nil {
		defer file.Close()
		os.MkdirAll("uploads/products", os.ModePerm)
		imageName := strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(handler.Filename)
		imagePath = "uploads/products/" + imageName
		dst, _ := os.Create(imagePath)
		defer dst.Close()
		io.Copy(dst, file)
	}

	tags := []string{}
	if tagsStr != "" {
		tags = strings.Split(tagsStr, ",")
	}

	product := model.Product{
		ID:          strconv.FormatInt(time.Now().UnixNano(), 10),
		Name:        name,
		Description: description,
		Price:       price,
		Image:       imagePath,
		Stock:       stock,
		SKU:         sku,
		Categories:  []string{categoryId},
		Tags:        tags,
		IsFeatured:  isFeatured,
		IsActive:    isActive,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	for i, cat := range categories {
		if cat.ID == categoryId {
			categories[i].Products = append(categories[i].Products, product)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}
