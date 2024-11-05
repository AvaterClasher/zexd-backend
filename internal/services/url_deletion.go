package services

import (
	"errors"

	"github.com/AvaterClasher/zexd/internal/daos"
	"github.com/AvaterClasher/zexd/internal/util"
)

func UrlDelete(shortenedUrl string) error {
	uid, encodedPart := util.GetUid(shortenedUrl)

	exists, err := rdb.Exists(ctx, encodedPart).Result()
	if err != nil {
		log.Errorf("Error checking Redis for URL existence: %s", err)
		return err
	}

	if exists == 0 {
		dbExists, err := daos.UrlExists(uid)
		if err != nil {
			return err
		}
		if !dbExists {
			return errors.New("URL does not exist")
		}
	}

	err = rdb.Del(ctx, encodedPart).Err()
	if err != nil {
		log.Errorf("Error in deletion from Redis: %s", err)
	}

	return daos.DeleteUrl(uid)
}
