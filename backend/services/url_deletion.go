package services

import (
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"strings"
	"zexd/daos"
)

func UrlDelete(shortenedUrl string) error {
	encodedPart := shortenedUrl[strings.LastIndex(shortenedUrl, "/")+1:]
	byteNumber, _ := base64.StdEncoding.DecodeString(encodedPart)
	uid, _ := strconv.Atoi(string(byteNumber))
	fmt.Println(uid)
	err := rdb.Del(ctx, shortenedUrl).Err()
	if err != nil {
		log.Println("Error in deletion from redis : ",err)
	}
	return daos.DeleteUrl(uid)
}
