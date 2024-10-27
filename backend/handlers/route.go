package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
)

func New() http.Handler {
	route := mux.NewRouter()
	route.HandleFunc("/check", BasicHandler)
	mainRouter := route.PathPrefix("/").Subrouter()
	mainRouter.HandleFunc("/api/create", CreateUrlHandler).Methods("POST", "OPTIONS")
	return route
}
