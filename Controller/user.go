package controller

import (
	"encoding/json"
	"fmt"
	"github/think.com/dots"
	"github/think.com/migler/index"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) CreateNewOrder(w http.ResponseWriter, r *http.Request) {
	var Order dots.Orders

	if err := json.NewDecoder(r.Body).Decode(&Order); err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid order data", err)
		return
	}
	defer r.Body.Close()
	fmt.Println(Order)
	if err := app.Storge.Post.CreateOrder(r.Context(), &Order, app.BleveSearchIndex); err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Failed to create order", err)
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "successfully",
		"message": "Order created successfully",
		"OrderID": Order.OrderID,
	})
}

// This section will handel the search in the program
func (app *Application) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	queryValue := chi.URLParam(r, "q")

	if app.BleveSearchIndex == nil {
		app.respondWithError(w, http.StatusInternalServerError, "BleveSearchIndex is not initialized", nil)
		return
	}

	res, err := index.BSearch(queryValue, app.BleveSearchIndex)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Search error", err)
		log.Printf("the err %s", err)
		log.Printf("the err %v", err)
		return
	}
	fmt.Println("this is from the search function\n")
	for title, cont := range res {
		log.Printf("the title is : %s the content is : %s\n", title, cont.(map[string]interface{})["short_description"])
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
