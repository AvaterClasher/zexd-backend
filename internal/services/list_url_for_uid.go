package services

import (
	"github.com/AvaterClasher/zexd/internal/daos"
)

type UrlData = daos.UrlData

func ListUrlForUid(uid string) ([]UrlData, error) {

	urls, err := daos.GetUrlsForUid(uid)

	if err != nil {
		log.Errorf("Error fetching URLs for user ID: %s", err)
		return nil, err
	}

	return urls, err
}