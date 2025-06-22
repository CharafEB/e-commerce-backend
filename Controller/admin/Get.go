package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

// GetImageHandler: gettin img's from product_img folder
func (app *Application) GetImageHandler(w http.ResponseWriter, r *http.Request) {
	imageName := chi.URLParam(r, "imageName")

	imagePath := filepath.Join("./products_img", imageName)
	fileInfo, err := os.Stat(imagePath)
	if os.IsNotExist(err) {
		app.respondWithError(w, http.StatusNotFound, "Image not found", err)
		return
	}

	ext := strings.ToLower(filepath.Ext(imageName))
	contentType := "application/octet-stream"
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".webp":
		contentType = "image/webp"
	case ".gif":
		contentType = "image/gif"
	case ".svg":
		contentType = "image/svg+xml"
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "public, max-age=31536000, immutable") // سنة + immutable
	w.Header().Set("Expires", time.Now().AddDate(1, 0, 0).Format(http.TimeFormat))
	w.Header().Set("Last-Modified", fileInfo.ModTime().Format(http.TimeFormat))
	w.Header().Set("ETag", fmt.Sprintf("\"%x\"", fileInfo.ModTime().UnixNano()))

	http.ServeContent(w, r, imageName, fileInfo.ModTime(), mustOpen(imagePath))
}

func mustOpen(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

// GetCategories : get the Categories from GetCategories table
func (app *Application) GetCategories(w http.ResponseWriter, r *http.Request) {
	Categorie, err := app.Storge.Get.GetCategories(r.Context())
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Error retrieving Categories:", err)
		return
	}
	fmt.Print(Categorie)

	if Categorie == nil {
		app.respondWithError(w, http.StatusNotFound, "There is no Categories pless add Categorie:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Categorie)

}

// GetSizes : get the Sizes from Size table
func (app *Application) GetSizes(w http.ResponseWriter, r *http.Request) {
	Sizes, err := app.Storge.Get.GetSizes(r.Context())
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Error retrieving Sizes:", err)
		return
	}

	if Sizes == nil {
		app.respondWithError(w, http.StatusNotFound, "There is no Sizes pless add Sizes:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Sizes)

}

// GetColors : get the Colors from Color table
func (app *Application) GetColors(w http.ResponseWriter, r *http.Request) {
	colors, err := app.Storge.Get.GetColors(r.Context())
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Error retrieving colors:", err)
		return
	}
	if colors == nil {
		app.respondWithError(w, http.StatusNotFound, "There is no colors pless add colors:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(colors)

}

// GetOrders : get the Orders
func (app *Application) GetOrders(w http.ResponseWriter, r *http.Request) {
	Orders, err := app.Storge.Orders.GetOrders(r.Context())
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Error retrieving Orders:", err)
		return
	}
	if Orders == nil {
		app.respondWithError(w, http.StatusNotFound, "There is no Orders pless add Orders:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Orders)

}

// GetOrdersByIs : get the Orders by ID
func (app *Application) GetOrdersById(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")

	if productID == "" {
		app.respondWithError(w, http.StatusInternalServerError, "Error Id is empty:", nil)

		return
	}
	Order, err := app.Storge.Orders.GetOrderByID(r.Context(), productID)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Error retrieving Order:", err)
		return
	}
	if Order == nil {
		app.respondWithError(w, http.StatusNotFound, "There is no Order pless add Order:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Order)

}
