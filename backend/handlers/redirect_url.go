package handlers

import (
	"log"
	"net/http"

	"zexd/services"
	"github.com/gorilla/mux"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortenedUrl := vars["shortenedUrl"]

	if shortenedUrl == "" {
		http.Error(w, "Shortened URL is required", http.StatusBadRequest)
		return
	}

	orgUrl, err := services.UrlRedirection(shortenedUrl)
	if err != nil || orgUrl == "" {
		log.Println("No matching URL found for redirection:", err)
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, orgUrl, http.StatusFound)
}