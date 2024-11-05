package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AvaterClasher/zexd/internal/services"
)

type inputListUrl struct {
	Url string `json:"url"`
}

type MetaUrlResponse struct {
	Metadata []services.ClickMetadata `json:"metadata"`
}

// ListMetaForLinksHandler godoc
// @Summary List Metadata for a specific Link
// @Description Returns a list of Metadata for the specified Link
// @Tags urls
// @operationId listMetadataForLinks
// @Accept  json
// @Produce  json
// @Param   input  body      inputListUrl  true  "URL to fetch metadata for"
// @Success 200     {object} MetaUrlResponse "Metadata for the specified URL"
// @Failure 400     {string} string "URL field is required"
// @Failure 404     {string} string "URL does not exist"
// @Failure 500     {string} string "Internal Server Error"
// @Router /api/list/url [post]
func ListMetaForLinksHandler(w http.ResponseWriter, r *http.Request) {
	var input inputListUrl
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

	metaData, err := services.ListMetaForLink(input.Url)
	if err != nil {
		log.Errorf("Error retrieving metadata for URL: %s", err)
		http.Error(w, "Could not find URL metadata", http.StatusInternalServerError)
		return
	}

	response := MetaUrlResponse{Metadata: metaData}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Errorf("Error encoding response: %s", err)
		http.Error(w, "Could not encode response", http.StatusInternalServerError)
	}
}
