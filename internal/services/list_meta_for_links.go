package services

import (
	"github.com/AvaterClasher/zexd/internal/daos"
	"github.com/AvaterClasher/zexd/internal/model"
	"github.com/AvaterClasher/zexd/internal/util"
)

type ClickMetadata = model.ClickMetadata

func ListMetaForLink(shortenedUrl string) ([]ClickMetadata, error) {
	uid, _ := util.GetUid(shortenedUrl)
	metaData, err := daos.GetClickMetadataForUrl(uid)

	if err != nil {
		log.Errorf("Error fetching Metadata for Link: %s", err)
		return nil, err
	}

	return metaData, err
}