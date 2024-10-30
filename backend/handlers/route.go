package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "zexd/docs"
)

type HistPrometheus struct {
	histogram *prometheus.HistogramVec
}

func (histProme *HistPrometheus) Populate() {
	histProme.histogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_durations_histogram_seconds",
		Help:    "HTTP request latency distributions.",
		Buckets: prometheus.ExponentialBuckets(0.0001, 1.5, 36),
	}, []string{"total_time_taken", "controller", "action"})

	prometheus.MustRegister(histProme.histogram)
}

func recordMetrics() {
	reqProcessed.Inc()
}

var (
	reqProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_requests",
		Help: "The total number of Incoming requests",
	})
)

func (histProme *HistPrometheus) PrometheusMonitoring(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			hist2 := histProme.histogram

			recordMetrics()
			t1 := time.Now()
			h.ServeHTTP(w, r)
			t2 := time.Now()

			hist2.WithLabelValues(
				FloatToString(t2.Sub(t1).Seconds()),
				r.URL.String(),
				r.Method).Observe(float64(time.Since(t1)) / float64(time.Second))
		})
}

func FloatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}

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
	hist.Populate()

	route.Handle("/api/metrics", promhttp.Handler()).Methods("GET")

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