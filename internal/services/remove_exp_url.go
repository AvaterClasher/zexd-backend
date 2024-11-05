package services

import (
	"os"
	"strconv"
	"time"

	"github.com/AvaterClasher/zexd/internal/daos"
	"github.com/joho/godotenv"
)

var exp_time int

func init() {
	exp_time = func() int {
		if expTimeStr, exists := os.LookupEnv("EXPIRY_TIME"); exists {
			if exp, err := strconv.Atoi(expTimeStr); err == nil {
				return exp
			}
		}
		_ = godotenv.Load(".env")
		if expTimeStr := os.Getenv("EXPIRY_TIME"); expTimeStr != "" {
			if exp, err := strconv.Atoi(expTimeStr); err == nil {
				return exp
			}
		}
		log.Warn("EXPIRY_TIME not set or invalid, defaulting to 60 minutes")
		return 60
	}()
}

func RemoveExpiredEntries() {
	expiryThreshold := time.Now().Add(-time.Duration(exp_time) * time.Minute)
	if err := daos.DeleteExpiredUrls(expiryThreshold); err != nil {
		log.Errorf("Error deleting expired URLs: %s", err)
	} else {
		log.Info("Expired URLs successfully deleted.")
	}
}
