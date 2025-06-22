package admin

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *Application) UpdatProduct(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	ids := queryParams["ids"]

	if len(ids) == 0 {
		log.Print("There is an err in the ids section :")
		http.Error(w, "ids are required", http.StatusBadRequest)
		return
	}

	if err := app.Storge.Update.UpdateEvent(r.Context(), ids); err != nil {
		http.Error(w, "article not found", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Events updated successfully",
		"ids":     ids,
	})

}

func (app *Application) ModifierProduct(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	ids := queryParams["ids"]

	if len(ids) == 0 {
		log.Print("There is an err in the ids section :")
		http.Error(w, "ids are required", http.StatusBadRequest)
		return
	}

	if err := app.Storge.Update.UpdateEvent(r.Context(), ids); err != nil {
		http.Error(w, "article not found", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  "success",
		"message": "Events updated successfully",
		"ids":     ids,
	})

}
