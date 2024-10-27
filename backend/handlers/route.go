package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
)

func New() http.Handler {
	route := mux.NewRouter()

	route.HandleFunc("/check", BasicHandler)

	return route
}
