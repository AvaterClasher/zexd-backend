package handlers

import (
	"fmt"
	"net/http"

	"zexd/services"

	"github.com/gorilla/mux"
)

// RedirectHandler godoc
// @Summary Redirect to the original URL from a shortened URL
// @Description Redirects to the original URL from a shortened URL code
// @Tags urls
// @operationId redirectUrl
// @Param   shortenedUrl  path     string  true  "Shortened URL code"
// @Success 302           {string} string  "Redirects to the original URL"
// @Failure 400           {string} string  "Shortened URL is required"
// @Failure 404           {string} string  "URL not found"
// @Router /{shortenedUrl} [get]
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortenedUrl := vars["shortenedUrl"]

	if shortenedUrl == "" {
		http.Error(w, "Shortened URL is required", http.StatusBadRequest)
		return
	}

	orgUrl, err := services.UrlRedirection(shortenedUrl)
	if err != nil || orgUrl == "" {
		log.Errorf("No matching URL found for redirection: %s", err)
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, orgUrl)
	// http.Redirect(w, r, orgUrl, http.StatusFound)
}