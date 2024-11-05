package main

import (
	"github.com/AvaterClasher/zexd/pkg/logger"
	"net/http"
	"github.com/AvaterClasher/zexd/internal/handlers"
	"github.com/AvaterClasher/zexd/internal/services"

	"github.com/jasonlvhit/gocron"
)

var log = logger.NewLogger()

func executeCronJob() {
	log.Infof("Executing Cron Service")
	gocron.Every(60).Minute().Do(services.RemoveExpiredEntries)
	<-gocron.Start()
}

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: handlers.New(),
	}

	go executeCronJob()
	log.Infof("Starting server on %s", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Errorf("%s", err)
	} else {
		log.Warnf("Server Closed")
	}
}