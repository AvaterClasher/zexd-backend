package daos

import (
	"context"
	"encoding/base64"
	"strconv"
	"time"
)

type UrlWithMetadata struct {
	Uid          int                 `json:"uid"`
	URL          string              `json:"url"`
	ShortenedUrl string              `json:"shortened_url"`
	Clicks       int                 `json:"clicks"`
	Metadata     []ClickWithMetadata `json:"metadata"`
}

type ClickWithMetadata struct {
	Id              int       `json:"id"`
	ClickedAt       time.Time `json:"clicked_at"`
	IpAddress       string    `json:"ip_address"`
	DeviceType      string    `json:"device_type"`
	OperatingSystem string    `json:"operating_system"`
	Referrer        string    `json:"referrer"`
	Browser         string    `json:"browser"`
}

func GetUrlsAndMetadataForUser(userID string) ([]UrlWithMetadata, error) {
	query := `
		SELECT 
			u.uid, u.url, u.clicks, 
			c.device_type, c.operating_system, c.referrer, c.browser, c.ip_address, c.clicked_at
		FROM 
			shortened_url u
		LEFT JOIN 
			click_metadata c 
		ON 
			u.uid = c.url_uid
		WHERE 
			u.user_id = $1
	`

	rows, err := db.Query(context.Background(), query, userID)
	if err != nil {
		log.Errorf("Error retrieving URLs and metadata for user: %v", err)
		return nil, err
	}
	defer rows.Close()

	urlMap := make(map[int]*UrlWithMetadata)

	for rows.Next() {
		var (
			uid             int
			url             string
			clicks          int
			deviceType      *string
			operatingSystem *string
			referrer        *string
			browser         *string
			ipAddress       *string
			clickedAt       *time.Time
		)

		if err := rows.Scan(
			&uid, &url, &clicks,
			&deviceType, &operatingSystem, &referrer, &browser, &ipAddress, &clickedAt,
		); err != nil {
			log.Errorf("Error scanning row: %v", err)
			return nil, err
		}

		byteNumber := []byte(strconv.Itoa(uid))
		shortenedUrl := Url + base64.StdEncoding.EncodeToString(byteNumber)

		urlData, exists := urlMap[uid]
		if !exists {
			urlData = &UrlWithMetadata{
				Uid:          uid,
				URL:          url,
				ShortenedUrl: shortenedUrl,
				Clicks:       clicks,
				Metadata:     []ClickWithMetadata{},
			}
			urlMap[uid] = urlData
		}
		metadataID := len(urlData.Metadata) + 1

		if deviceType != nil && operatingSystem != nil {
			urlData.Metadata = append(urlData.Metadata, ClickWithMetadata{
				Id:              metadataID,
				DeviceType:      *deviceType,
				OperatingSystem: *operatingSystem,
				Referrer:        *referrer,
				Browser:         *browser,
				IpAddress:       *ipAddress,
				ClickedAt:       *clickedAt,
			})
		}
	}

	if err := rows.Err(); err != nil {
		log.Errorf("Error iterating through rows: %v", err)
		return nil, err
	}

	urlsWithMetadata := make([]UrlWithMetadata, 0, len(urlMap))
	for _, urlData := range urlMap {
		urlsWithMetadata = append(urlsWithMetadata, *urlData)
	}

	return urlsWithMetadata, nil
}
