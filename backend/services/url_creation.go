package services

import (
	"encoding/base64"
	"os"
	"strconv"
	"zexd/daos"
	"zexd/model"
)

var url string
var set bool

func init() {
	url, set = os.LookupEnv("SERVER_DOMAIN")
	if !set {
		url = "http://localhost:8080/"
	}
}

func CreateShortenedUrl(inputUrl string, user_id string) string {
	uid, err := daos.InsertShortenedUrl(model.UrlModel{Url: inputUrl, User_id: user_id})
	if err != nil {
		return ""
	}
	byteNumber := []byte(strconv.Itoa(uid))
	tempUrl := base64.StdEncoding.EncodeToString(byteNumber)
	newUrl := url + tempUrl
	
	return newUrl
}