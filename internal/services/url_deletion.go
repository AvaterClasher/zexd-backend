package services

import (
	"encoding/base64"
	"errors"
	"strconv"
	"strings"
	"github.com/AvaterClasher/zexd/internal/daos"
)

func UrlDelete(shortenedUrl string) error {
	encodedPart := shortenedUrl[strings.LastIndex(shortenedUrl, "/")+1:]
	
	byteNumber, _ := base64.StdEncoding.DecodeString(encodedPart)
	uid, _ := strconv.Atoi(string(byteNumber))

	exists, err := rdb.Exists(ctx, encodedPart).Result()
	if err != nil {
		log.Errorf("Error checking Redis for URL existence: %s", err)
		return err
	}

	if exists == 0 {
		dbExists, err := daos.UrlExists(uid)
		if err != nil {
			return err
		}
		if !dbExists {
			return errors.New("URL does not exist")
		}
	}

	err = rdb.Del(ctx, encodedPart).Err()
	if err != nil {
		log.Errorf("Error in deletion from Redis: %s", err)
	}

	return daos.DeleteUrl(uid)
}
