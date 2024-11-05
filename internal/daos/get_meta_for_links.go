package daos

import (
	"context"

	"github.com/AvaterClasher/zexd/internal/model"
)

func GetClickMetadataForUrl(urlUid int) ([]model.ClickMetadata, error) {
	var metadata []model.ClickMetadata

	query := `
		SELECT id, device_type, operating_system, referrer, browser, ip_address, clicked_at 
		FROM click_metadata 
		WHERE url_uid = $1
	`
	rows, err := db.Query(context.Background(), query, urlUid)
	if err != nil {
		log.Errorf("Error retrieving click metadata for URL uid %d: %v", urlUid, err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var clickMeta model.ClickMetadata
		if err := rows.Scan(
			&clickMeta.ID,
			&clickMeta.DeviceType,
			&clickMeta.OperatingSystem,
			&clickMeta.Referrer,
			&clickMeta.Browser,
			&clickMeta.IpAddress,
			&clickMeta.ClickedAt,
		); err != nil {
			log.Errorf("Error scanning row for click metadata: %v", err)
			return nil, err
		}

		metadata = append(metadata, clickMeta)
	}

	if err := rows.Err(); err != nil {
		log.Errorf("Error iterating through click metadata rows: %v", err)
		return nil, err
	}

	return metadata, nil
}
