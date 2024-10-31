package services

import (
	"os"
	"strconv"
	"time"

	"zexd/daos"
)

var exp_time int

func init() {
	expTimeStr, exists := os.LookupEnv("EXPIRY_TIME")
	if exists {
		if parsedExpTime, err := strconv.Atoi(expTimeStr); err == nil {
			exp_time = parsedExpTime
		} else {
			log.Error("Invalid EXPIRY_TIME value, defaulting to 60 minutes")
			exp_time = 60
		}
	} else {
		log.Error("EXPIRY_TIME not set, defaulting to 60 minutes")
		exp_time = 60
	}
}

func RemoveExpiredEntries() {
	expiryThreshold := time.Now().Add(-time.Duration(exp_time) * time.Minute)
	if err := daos.DeleteExpiredUrls(expiryThreshold); err != nil {
		log.Errorf("Error deleting expired URLs: %s", err)
	} else {
		log.Error("Expired URLs successfully deleted.")
	}
}