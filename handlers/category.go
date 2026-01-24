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

var categories []model.Category

// ---------------- Admin ----------------

// Create Category with image support
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	contentType := r.Header.Get("Content-Type")
	var cat model.Category

	if strings.HasPrefix(contentType, "multipart/form-data") {
		// Multipart form for file upload
		r.ParseMultipartForm(10 << 20)

		cat.Name = r.FormValue("name")
		imageFile, handler, err := r.FormFile("image")
		if err == nil {
			defer imageFile.Close()
			os.MkdirAll("uploads/categories", os.ModePerm)
			imageName := strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(handler.Filename)
			imagePath := "uploads/categories/" + imageName
			dst, _ := os.Create(imagePath)
			defer dst.Close()
			io.Copy(dst, imageFile)
			cat.Image = imagePath
		} else {
			// if file not uploaded, can use URL
			cat.Image = r.FormValue("image_url")
		}
	} else {
		// JSON request
		json.NewDecoder(r.Body).Decode(&cat)
	}

	if cat.Name == "" {
		http.Error(w, "Category name required", http.StatusBadRequest)
		return
	}

	cat.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	cat.Products = []model.Product{}
	categories = append(categories, cat)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cat)
}

// Update Category with image support
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	var updateCat model.Category
	contentType := r.Header.Get("Content-Type")

	if strings.HasPrefix(contentType, "multipart/form-data") {
		r.ParseMultipartForm(10 << 20)
		updateCat.Name = r.FormValue("name")
		file, handler, err := r.FormFile("image")
		if err == nil {
			defer file.Close()
			os.MkdirAll("uploads/categories", os.ModePerm)
			imageName := strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(handler.Filename)
			imagePath := "uploads/categories/" + imageName
			dst, _ := os.Create(imagePath)
			defer dst.Close()
			io.Copy(dst, file)
			updateCat.Image = imagePath
		} else {
			updateCat.Image = r.FormValue("image_url")
		}
	} else {
		json.NewDecoder(r.Body).Decode(&updateCat)
	}

	for i, cat := range categories {
		if cat.ID == id {
			if updateCat.Name != "" {
				categories[i].Name = updateCat.Name
			}
			if updateCat.Image != "" {
				categories[i].Image = updateCat.Image
			}
			json.NewEncoder(w).Encode(categories[i])
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

// Delete Category
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	for i, cat := range categories {
		if cat.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

// ---------------- End User ----------------

// Get all categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

// Get category by ID
func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	for _, cat := range categories {
		if cat.ID == id {
			json.NewEncoder(w).Encode(cat)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}
