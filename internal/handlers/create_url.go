package handlers

import (
	"encoding/json"
	"io"
	"github.com/AvaterClasher/zexd/pkg/logger"
	"net/http"

	"github.com/AvaterClasher/zexd/internal/services"
)

var log = logger.NewLogger()

type inputUrl struct {
	Url     string `json:"url"`
	User_id string `json:"user_id"`
}

type shortenedUrlResponse struct {
	ShortenedUrl string `json:"shortened_url"`
}

// CreateUrlHandler godoc
// @Summary Shorten a URL
// @Description Creates a shortened version of the provided URL
// @Tags urls
// @operationId createShortenedUrl
// @Accept  json
// @Produce  json
// @Param   input  body      inputUrl  true  "URL and User ID"
// @Success 200     {object} shortenedUrlResponse
// @Failure 400     {string} string "Invalid JSON format or missing fields"
// @Failure 500     {string} string "Internal Server Error"
// @Router /api/create [post]
func CreateUrlHandler(w http.ResponseWriter, r *http.Request) {
	var input inputUrl
	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		log.Error("Encountered an error while reading request body")
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(reqBody, &input); err != nil {
		log.Error("Error unmarshalling JSON", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if input.Url == "" {
		log.Error("Url is empty")
		http.Error(w, "Url is empty", http.StatusBadRequest)
		return
	}

	if input.User_id == "" {
		log.Error("User_id is empty")
		http.Error(w, "User_id is empty", http.StatusBadRequest)
		return
	}

	shortenedUrl := services.CreateShortenedUrl(input.Url, input.User_id)
	response := shortenedUrlResponse{ShortenedUrl: shortenedUrl}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Errorf("Error encoding response: %s", err)
		http.Error(w, "Could not encode response", http.StatusInternalServerError)
	}
}
