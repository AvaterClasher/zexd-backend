package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"zexd/services"

	"github.com/gorilla/mux"
)

func ListUrlsForUidHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]

	if user_id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	urls, err := services.ListUrlForUid(user_id)
	if err != nil {
		log.Println("Error fetching URLs for user ID:", err)
		http.Error(w, "Error fetching URLs for user ID", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(urls); err != nil {
		log.Println("Error encoding response", err)
		http.Error(w, "Could not encode response", http.StatusInternalServerError)
	}
}