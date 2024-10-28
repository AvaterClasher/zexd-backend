package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"zexd/services"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var input inputUrl
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if input.Url == "" {
		log.Println("URL field is empty in request body")
		http.Error(w, "URL field is required", http.StatusBadRequest)
		return
	}

	if err := services.UrlDelete(input.Url); err != nil {
		log.Println("Error deleting URL:", err)
		http.Error(w, "Could not delete the URL", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "URL successfully deleted"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error encoding JSON response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
