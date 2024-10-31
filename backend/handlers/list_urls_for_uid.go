package handlers

import (
	"encoding/json"
	"net/http"

	"zexd/services"

	"github.com/gorilla/mux"
)

// ListUrlsForUidHandler godoc
// @Summary List URLs for a specific user ID
// @Produce  json
// @Param   user_id  path     string  true  "User ID"
// @Success 200      {array}  string  "List of URLs for the user"
// @Failure 400      {string} string  "User ID is required"
// @Failure 500      {string} string  "Internal Server Error"
// @Router /api/list/{user_id} [get]
func ListUrlsForUidHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user_id := vars["user_id"]

	if user_id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	urls, err := services.ListUrlForUid(user_id)
	if err != nil {
		log.Errorf("Error fetching URLs for user ID: %s", err)
		http.Error(w, "Error fetching URLs for user ID", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(urls); err != nil {
		log.Errorf("Error encoding response: %s", err)
		http.Error(w, "Could not encode response", http.StatusInternalServerError)
	}
}