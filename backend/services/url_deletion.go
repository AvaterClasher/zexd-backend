package services

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"zexd/daos"
)

func UrlDelete(shortenedUrl string) error {
	encodedPart := shortenedUrl[strings.LastIndex(shortenedUrl, "/")+1:]
	byteNumber, _ := base64.StdEncoding.DecodeString(encodedPart)
	uid, _ := strconv.Atoi(string(byteNumber))
	fmt.Println(uid)
	return daos.DeleteUrl(uid)
}
