package admin

import (
	"encoding/json"
	"net/http"
)

func (app *Application) respondWithError(w http.ResponseWriter, code int, message string, err error) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "invalid",
		"message": message,
		"problem": err.Error(),
	})
}
