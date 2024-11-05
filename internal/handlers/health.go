package handlers

import (
	"fmt"
	"net/http"
)

// HealthCheckHandler godoc
// @Summary Check if the server is online
// @Description Returns a message indicating the server is online
// @Tags health
// @operationId healthCheck
// @Produce text/plain
// @Success 200 {string} string "Server is online"
// @Router /health [get]
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Server is online")
}
