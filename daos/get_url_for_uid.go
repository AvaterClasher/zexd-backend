package daos

import (
	"encoding/base64"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type UrlData struct {
	Uid          int    `json:"uid"`
	URL          string `json:"url"`
	ShortenedUrl string `json:"shortened_url"`
	Clicks       int    `json:"clicks"`
}

var Url string
var Set bool

func init() {
	Url, Set = os.LookupEnv("SERVER_DOMAIN")
	if !Set {
		Url = "http://localhost:8080/"
	}
}

func GetUrlsForUid(user_id string) ([]UrlData, error) {
	var urls []UrlData

	query := `
		SELECT uid, url, clicks 
		FROM shortened_url 
		WHERE user_id = $1`
	rows, err := db.Query(query, user_id)
	if err != nil {
		log.Errorf("Error retrieving URLs for user: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var urlData UrlData
		if err := rows.Scan(&urlData.Uid, &urlData.URL, &urlData.Clicks); err != nil {
			log.Errorf("Error scanning row: %v", err)
			return nil, err
		}

		byteNumber := []byte(strconv.Itoa(urlData.Uid))
		urlData.ShortenedUrl = Url + base64.StdEncoding.EncodeToString(byteNumber)

		urls = append(urls, urlData)
	}

	if err := rows.Err(); err != nil {
		log.Errorf("Error retrieving rows: %v", err)
		return nil, err
	}

	return urls, nil
}
