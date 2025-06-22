package admin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github/think.com/dots"

	"time"
)

func (app *Application) CreatNewProduct(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var event dots.Product

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid JSON format", err)
		return
	}

	if err := app.Storge.Post.CreateProducts(r.Context(), &event, app.BleveSearchIndex); err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Event did not created successfully", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":          "success",
		"message":         "Event created successfully",
		"ProductsName : ": event.ProductsName,
	})
}

// CreatNewCategorie: Create a new Categorie to hte categorie table
func (app *Application) CreatNewCategorie(w http.ResponseWriter, r *http.Request) {

	var Categorie dots.Categories

	if err := json.NewDecoder(r.Body).Decode(&Categorie); err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid JSON format", err)
		return
	}

	defer r.Body.Close()

	if err := app.Storge.Post.AddCategories(r.Context(), Categorie.CategorieValue); err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Categorie did not created successfully ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":          "success",
		"message":         "Categorie Value added successfully",
		"Categorie value": Categorie.CategorieValue,
	})
}

// CreatNewSize: Create a new Size to hte Size table
func (app *Application) CreatNewSize(w http.ResponseWriter, r *http.Request) {

	var Size dots.Size

	if err := json.NewDecoder(r.Body).Decode(&Size); err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid JSON format", err)
		return
	}

	defer r.Body.Close()

	if err := app.Storge.Post.AddSize(r.Context(), Size.SizeValue); err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Categorie did not created successfully ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":     "success",
		"message":    "Size Value added successfully",
		"Size Value": Size.SizeValue,
	})
}

// CreatNewColor: Create a new Color to hte Color table
func (app *Application) CreatNewColor(w http.ResponseWriter, r *http.Request) {

	var Color dots.Color

	if err := json.NewDecoder(r.Body).Decode(&Color); err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid JSON format", err)
		return
	}

	defer r.Body.Close()

	if err := app.Storge.Post.AddColor(r.Context(), Color.ColorValue); err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Categorie did nit created successfully ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":      "success",
		"message":     "Color Value added successfully",
		"Color Value": Color.ColorValue,
	})
}

// UploadProductsImgs: Upload New img for Products
func (app *Application) UploadProductsImgs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	const (
		maxUploadSize = 4 * 1024 * 1024 //4 bm as max
		maxFiles      = 4               // number of pic's
	)

	if err := r.ParseMultipartForm(maxUploadSize * maxFiles); err != nil {
		fmt.Println(err)
		http.Error(w, "Total files size too big", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["images"]

	if len(files) > maxFiles {
		http.Error(w, fmt.Sprintf("Maximum %d images allowed", maxFiles), http.StatusBadRequest)
		return
	}

	uploadDir := "./products_img"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		fmt.Println(err)
		http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
		return
	}

	var uploadedFiles []string

	for _, fileHeader := range files {

		ext := filepath.Ext(fileHeader.Filename)
		switch strings.ToLower(ext) {
		case ".jpg", ".jpeg", ".png":
		default:
			http.Error(w, fmt.Sprintf("Unsupported file type: %s", ext), http.StatusBadRequest)
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid image", http.StatusBadRequest)
			return
		}
		defer file.Close()

		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), fileHeader.Filename)
		dstPath := filepath.Join(uploadDir, filename)

		dst, err := os.Create(dstPath)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error while creating the file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			fmt.Println(err)
			http.Error(w, "Error while saving the file", http.StatusInternalServerError)
			return
		}

		uploadedFiles = append(uploadedFiles, filename)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":   "success",
		"fileURLs": uploadedFiles,
		"message":  fmt.Sprintf("%d files uploaded successfully", len(uploadedFiles)),
	})
}
