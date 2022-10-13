package handlers

import (
	"api-postgresql/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error decoding todo: %v", err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error": true,
			"Message": fmt.Sprintf("Error creating todo: %v", err),
		}
	}else {
		resp = map[string]any{
            "Error": false,
			"Message": fmt.Sprintf("Todo created successfully! ID: %d", id),
        }
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
