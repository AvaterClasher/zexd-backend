package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"zexd/services"
)

type inputUrl struct {
	Url     string `json:"url"`
	User_id string `json:"user_id"`
}

type shortenedUrlResponse struct {
	ShortenedUrl string `json:"shortened_url"`
}

func CreateUrlHandler(w http.ResponseWriter, r *http.Request) {
	var input inputUrl
	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println("Encountered an error while reading request body")
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(reqBody, &input); err != nil {
		log.Println("Error unmarshalling JSON", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if input.Url == "" {
		log.Println("Url is empty")
		http.Error(w, "Url is empty", http.StatusBadRequest)
		return
	}

	if input.User_id == "" {
		log.Println("User_id is empty")
		http.Error(w, "User_id is empty", http.StatusBadRequest)
		return
	}

	shortenedUrl := services.CreateShortenedUrl(input.Url, input.User_id)
	response := shortenedUrlResponse{ShortenedUrl: shortenedUrl}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error encoding response", err)
		http.Error(w, "Could not encode response", http.StatusInternalServerError)
	}
}
