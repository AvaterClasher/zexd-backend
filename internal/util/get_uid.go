package util

import (
	"encoding/base64"
	"strconv"
	"strings"
)

func GetUid(shortenedUrl string) (int, string) {
	encodedPart := shortenedUrl[strings.LastIndex(shortenedUrl, "/")+1:]
	byteNumber, _ := base64.StdEncoding.DecodeString(encodedPart)
	uid, _ := strconv.Atoi(string(byteNumber))

	return uid, encodedPart
}