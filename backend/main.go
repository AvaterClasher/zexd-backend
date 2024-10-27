package main

import (
	"log"
	"net/http"
	"zexd/handlers"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: handlers.New(),
	}

	log.Printf("Starting server on %s", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%s", err)
	} else {
		log.Printf("Server Closed")
	}
}