package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
)

func New() http.Handler {
	route := mux.NewRouter()
	route.HandleFunc("/check", BasicHandler)
	mainRouter := route.PathPrefix("/").Subrouter()

	mainRouter.HandleFunc("/{shortenedUrl}", RedirectHandler)
	mainRouter.HandleFunc("/api/create", CreateUrlHandler).Methods("POST", "OPTIONS")
	mainRouter.HandleFunc("/api/delete", DeleteHandler).Methods("POST", "OPTIONS")
	mainRouter.HandleFunc("/api/list/{user_id}", ListUrlsForUidHandler).Methods("GET", "OPTIONS")

	return route
}