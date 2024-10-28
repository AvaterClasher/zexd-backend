package services

import (
	"log"
	"zexd/daos"
)

type UrlData = daos.UrlData

func ListUrlForUid(uid string) ([]UrlData, error) {

	urls, err := daos.GetUrlsForUid(uid)

	if err != nil {
		log.Println("Error in getting URL from database", err)
		return nil, err
	}

	return urls, err
}