package services

import (
	"github.com/AvaterClasher/zexd/internal/daos"
)

func ListUrlsWithMetadata(userID string) ([]daos.UrlWithMetadata, error) {
	urlsWithMetadata, err := daos.GetUrlsAndMetadataForUser(userID)
	if err != nil {
		log.Errorf("Error fetching URLs and metadata for user %s: %v", userID, err)
		return nil, err
	}
	return urlsWithMetadata, nil
}
