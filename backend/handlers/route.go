package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "zexd/docs"
	middleware "zexd/middleware"
)

type HistPrometheus struct {
	histogram *prometheus.HistogramVec
}

func (histProme *HistPrometheus) PrometheusMonitoring(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			recordMetrics()
			incrementMethodCounter(r.Method)
			h.ServeHTTP(w, r)
		})
}

func recordMetrics() {
	reqProcessed.Inc()
}

func incrementMethodCounter(method string) {
	switch method {
	case http.MethodGet:
		reqProcessedGET.Inc()
	case http.MethodPost:
		reqProcessedPOST.Inc()
	case http.MethodDelete:
		reqProcessedDELETE.Inc()
	}
}

var (
	reqProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_requests",
		Help: "The total number of incoming requests",
	})
	reqProcessedGET = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_get_requests",
		Help: "The total number of GET requests",
	})
	reqProcessedPOST = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_post_requests",
		Help: "The total number of POST requests",
	})
	reqProcessedDELETE = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_delete_requests",
		Help: "The total number of DELETE requests",
	})
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

	hist := HistPrometheus{}

	route.Handle("/api/metrics", promhttp.Handler()).Methods("GET")

	route.HandleFunc("/health", HealthCheckHandler)

	mainRouter := route.PathPrefix("/").Subrouter()
	mainRouter.Use(middleware.CORSMiddleware)
	mainRouter.Use(hist.PrometheusMonitoring)
	mainRouter.Use(middleware.LogMiddleware)

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
