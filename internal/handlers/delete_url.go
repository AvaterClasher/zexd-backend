package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"github.com/AvaterClasher/zexd/internal/services"
)

type inputDelUrl struct {
	Url string `json:"url"` // URL to delete
}


// DeleteHandler godoc
// @Summary Delete a shortened URL
// @Description Deletes a shortened URL
// @Tags urls
// @operationId deleteShortenedUrl
// @Accept  json
// @Produce  json
// @Param   input  body      inputDelUrl  true  "URL to delete"
// @Success 200     {object} map[string]string "Success message"
// @Failure 400     {string} string "URL field is required"
// @Failure 404     {string} string "URL does not exist"
// @Failure 500     {string} string "Internal Server Error"
// @Router /api/delete [post]
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var input inputDelUrl
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Errorf("Error decoding request body: %s", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if input.Url == "" {
		log.Error("URL field is empty in request body")
		http.Error(w, "URL field is required", http.StatusBadRequest)
		return
	}

	err := services.UrlDelete(input.Url)
	if err != nil {
		if errors.Is(err, errors.New("URL does not exist")) {
			log.Errorf("URL does not exist: %s", input.Url)
			http.Error(w, "URL does not exist", http.StatusNotFound)
		} else {
			log.Errorf("Error deleting URL: %s", err)
			http.Error(w, "Could not delete the URL", http.StatusInternalServerError)
		}
		return
	}

	response := map[string]string{"message": "URL successfully deleted"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Errorf("Error encoding JSON response: %s", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
