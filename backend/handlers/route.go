package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "zexd/docs"
)

// Swagger Info
// @title ZexD API
// @version 1.0
// @description API for a URL Shortener Service.

// @license.name MIT
// @license.url https://opensource.org/license/mit

// @host localhost:8080
// @BasePath /
func New() http.Handler {
	route := mux.NewRouter()
	route.HandleFunc("/health", HealthCheckHandler)
	mainRouter := route.PathPrefix("/").Subrouter()

	mainRouter.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	mainRouter.HandleFunc("/{shortenedUrl}", RedirectHandler)
	mainRouter.HandleFunc("/api/create", CreateUrlHandler).Methods("POST", "OPTIONS")
	mainRouter.HandleFunc("/api/delete", DeleteHandler).Methods("POST", "OPTIONS")
	mainRouter.HandleFunc("/api/list/{user_id}", ListUrlsForUidHandler).Methods("GET", "OPTIONS")

	return route
}