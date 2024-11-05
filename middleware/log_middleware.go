package middleware

import (
	"net/http"
	"github.com/AvaterClasher/zexd/logger"
)

var log = logger.NewLoggerforHttp()

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}