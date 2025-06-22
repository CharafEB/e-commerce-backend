package controller

import (
	"encoding/json"
	webrander "github/think.com/migler/WebRander"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) GetContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	queryValue := chi.URLParam(r, "q")

	if queryValue == "" {
		app.respondWithError(w, http.StatusBadRequest, "id is required", nil)
		return
	}
	res, err := app.Storge.Get.GetContentSearch(r.Context(), queryValue)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "article not found", err)
		return
	}

	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	webrander.ContentR(w, res, queryValue)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
