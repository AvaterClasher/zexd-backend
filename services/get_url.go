package services

import (
	"encoding/base64"
	"github.com/AvaterClasher/zexd/logger"
	"strconv"
	"strings"
	"github.com/AvaterClasher/zexd/daos"

	"github.com/redis/go-redis/v9"
)

var log = logger.NewLogger()

func UrlRedirection(shortenedUrl string) (string, error) {
	inputUrl, err := rdb.Get(ctx, shortenedUrl).Result()
	if err == nil {
		// log.Infof("Shortened URL found in cache: %s", inputUrl)
		encodedPart := shortenedUrl[strings.LastIndex(shortenedUrl, "/")+1:]
		byteNumber, _ := base64.StdEncoding.DecodeString(encodedPart)
		uid, _ := strconv.Atoi(string(byteNumber))
		go func() {
			if dbErr := daos.IncrementClickCount(uid); dbErr != nil {
				log.Errorf("Error incrementing click count: %s", dbErr)
			}
		}()
		return inputUrl, nil
	} else if err != redis.Nil {
		log.Errorf("Redis error: %s", err)
	}

	encodedPart := shortenedUrl[strings.LastIndex(shortenedUrl, "/")+1:]
	byteNumber, _ := base64.StdEncoding.DecodeString(encodedPart)
	uid, _ := strconv.Atoi(string(byteNumber))

	inputUrl, err = daos.GetUrlAndIncrement(uid)
	if err != nil {
		log.Errorf("Shortened Url: %s", url + shortenedUrl)
		return "", err
	}

	return inputUrl, err
}