package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AvaterClasher/zexd/internal/services"
	"github.com/gorilla/mux"
)

// ListUrlsWithMetadataHandler godoc
// @Summary List all URLs and metadata for a user
// @Description Returns all URLs and their associated metadata for a specific user ID
// @Tags urls
// @operationId listUrlsWithMetadata
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {array} daos.UrlWithMetadata "List of URLs with metadata"
// @Failure 400 {string} string "User ID is required"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/list/{user_id}/all [get]
func ListUrlsWithMetadataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	urlsWithMetadata, err := services.ListUrlsWithMetadata(userID)
	if err != nil {
		log.Errorf("Error fetching URLs and metadata for user %s: %v", userID, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(urlsWithMetadata); err != nil {
		log.Errorf("Error encoding response: %s", err)
		http.Error(w, "Could not encode response", http.StatusInternalServerError)
	}
}
