package main

import (
	"log"
	"net/http"
	"zexd/handlers"
	"zexd/services"

	"github.com/jasonlvhit/gocron"
)

func executeCronJob() {
	log.Println("**Executing Cron Service**")
	gocron.Every(60).Minute().Do(services.RemoveExpiredEntries)
	<-gocron.Start()
}

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: handlers.New(),
	}

	go executeCronJob()

	log.Printf("Starting server on %s", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%s", err)
	} else {
		log.Printf("Server Closed")
	}
}