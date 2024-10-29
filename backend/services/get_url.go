package services

import (
	"encoding/base64"
	"log"
	"strconv"
	"strings"
	"zexd/daos"

	"github.com/redis/go-redis/v9"
)

func UrlRedirection(shortenedUrl string) (string, error) {
	inputUrl, err := rdb.Get(ctx, shortenedUrl).Result()
	if err == nil {
		log.Println("Shortened URL found in cache", inputUrl)
		encodedPart := shortenedUrl[strings.LastIndex(shortenedUrl, "/")+1:]
		byteNumber, _ := base64.StdEncoding.DecodeString(encodedPart)
		uid, _ := strconv.Atoi(string(byteNumber))
		go func() {
			if dbErr := daos.IncrementClickCount(uid); dbErr != nil {
				log.Println("Error incrementing click count:", dbErr)
			}
		}()
		return inputUrl, nil
	} else if err != redis.Nil {
		log.Println("Redis error:", err)
	}

	encodedPart := shortenedUrl[strings.LastIndex(shortenedUrl, "/")+1:]
	byteNumber, _ := base64.StdEncoding.DecodeString(encodedPart)
	uid, _ := strconv.Atoi(string(byteNumber))

	inputUrl, err = daos.GetUrlAndIncrement(uid)
	log.Println("URL from database", inputUrl)
	if err != nil {
		log.Println("Error in getting URL from database", err)
		return "", err
	}

	return inputUrl, err
}