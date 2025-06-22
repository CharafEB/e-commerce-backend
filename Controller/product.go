package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// This section will handel the getting of the product content
func (app *Application) GetProductContent(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")

	if productID == "" {
		app.respondWithError(w, http.StatusBadRequest, "Product ID is required", nil)
		return
	}

	product, err := app.Storge.Get.GetProductContent(r.Context(), productID)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Error retrieving product content", err)
		return
	}

	if product == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// This section will handel the getting of the product by category.

func (app *Application) GetProductsByCategory(w http.ResponseWriter, r *http.Request) {

	category := chi.URLParam(r, "category")

	if category == "" {
		app.respondWithError(w, http.StatusBadRequest, "Category is required", nil)
		return
	}

	products, err := app.Storge.Get.GetByCategorie(r.Context(), category)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Error retrieving products by category", err)
		return
	}

	if len(products) == 0 {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// This section will handel the getting of the product by price under the specified value.

func (app *Application) GetProductsByUPrice(w http.ResponseWriter, r *http.Request) {

	UPrice := chi.URLParam(r, "price")

	UPriceValur, err := strconv.ParseFloat(UPrice, 64)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid price format", err)
		return
	}

	if UPriceValur == 0 || UPriceValur < 0 {
		app.respondWithError(w, http.StatusBadRequest, "UPriceValur is required", nil)
		return
	}

	products, err := app.Storge.Get.GetByPriceUnder(r.Context(), UPriceValur)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Error retrieving products by price", err)
		return
	}

	if len(products) == 0 {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// This section will handel the getting of the product by price over the specified value.

func (app *Application) GetProductsByOPrice(w http.ResponseWriter, r *http.Request) {

	UPrice := chi.URLParam(r, "price")

	UPriceValur, err := strconv.ParseFloat(UPrice, 64)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid price format", err)
		return
	}

	if UPriceValur == 0 || UPriceValur < 0 {
		app.respondWithError(w, http.StatusBadRequest, "UPriceValur is required", nil)
		return
	}

	products, err := app.Storge.Get.GetByPriceUper(r.Context(), UPriceValur)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Error retrieving products by price", err)
		return
	}

	if len(products) == 0 {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// This section will handel the getting of the product by name.
func (app *Application) GetAllProducts(w http.ResponseWriter, r *http.Request) {

	products, err := app.Storge.Get.GetProducts(r.Context())
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Error retrieving all products", err)
		return
	}

	if len(products) == 0 {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
