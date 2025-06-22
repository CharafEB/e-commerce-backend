package admin

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)
//DeleteProduct:
func (app *Application) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "q")

	if id == "" {
		log.Print("There is an err in the id section :")
		http.Error(w, "ids are required", http.StatusBadRequest)
		return
	}

	if err := app.Storge.Delete.DeleteProducts(r.Context(), id); err != nil {

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Events has been deleted successfully",
		"id":      id,
	})

}
//DeleteCategories:
func (app *Application) DeleteCategories(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "q")

	if id == "" {
		log.Print("There is an err in the id section :")
		http.Error(w, "ids are required", http.StatusBadRequest)
		return
	}

	if err := app.Storge.Delete.DeleteCategories(r.Context(), id); err != nil {

		app.respondWithError(w, http.StatusInternalServerError, "Categories not deleted:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Categorie has been deleted successfully",
		"id":      id,
	})

}
//DeleteColor:
func (app *Application) DeleteColor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "q")

	if id == "" {
		log.Print("There is an err in the id section :")
		http.Error(w, "ids are required", http.StatusBadRequest)
		return
	}

	if err := app.Storge.Delete.DeleteColor(r.Context(), id); err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "color not deleted:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "color has been deleted successfully",
		"id":      id,
	})

}
//DeleteSize:
func (app *Application) DeleteSize(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "q")

	if id == "" {
		log.Print("There is an err in the id section :")
		http.Error(w, "ids are required", http.StatusBadRequest)
		return
	}

	if err := app.Storge.Delete.DeleteSize(r.Context(), id); err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "size not deleted:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Size has been deleted successfully",
		"id":      id,
	})

}
//DeleteOrder:
func (app *Application) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "q")

	if id == "" {
		log.Print("There is an err in the id section :")
		http.Error(w, "ids are required", http.StatusBadRequest)
		
		return
	}

	if err := app.Storge.Delete.DeleteProducts(r.Context(), id); err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "Order not Deleted:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "order has been deleted successfully",
		"id":      id,
	})

}
