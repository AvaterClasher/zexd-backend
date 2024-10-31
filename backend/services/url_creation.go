package services

import (
	"encoding/base64"
	"os"
	"strconv"
	"time"
	"zexd/cache"
	"zexd/daos"
	"zexd/model"

	"github.com/redis/go-redis/v9"
)

var url string
var set bool

var ctx = cache.Ctx

var rdb *redis.Client

func init() {
	rdb = cache.CreateCon()
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

	err = rdb.Set(ctx, tempUrl, inputUrl, time.Duration(exp_time)*time.Minute).Err()
	if err != nil {
		log.Printf("Error in setting Redis value: %v", err)
	}

	return newUrl
}
