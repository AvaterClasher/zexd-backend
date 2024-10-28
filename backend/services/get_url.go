package services

import (
	"encoding/base64"
	"log"
	"strconv"
	"strings"
	"zexd/daos"
)

func UrlRedirection(shortenedUrl string) (string, error) {

	encodedPart := shortenedUrl[strings.LastIndex(shortenedUrl, "/")+1:]
	byteNumber, _ := base64.StdEncoding.DecodeString(encodedPart)
	uid, _ := strconv.Atoi(string(byteNumber))

	inputUrl, err := daos.GetUrlAndIncrement(uid)
	if err != nil {
		log.Println("Error in getting URL from database", err)
		return "", err
	}

	return inputUrl, err
}
