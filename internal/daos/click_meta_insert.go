package daos

import (
	"context"

	"github.com/AvaterClasher/zexd/internal/model"
)

func InsertClickMetadata(clickData model.ClickMetadata) error {
	query := `INSERT INTO click_metadata (url_uid, clicked_at, ip_address, device_type, operating_system, referrer, browser) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(context.Background(), query, clickData.UrlUid, clickData.ClickedAt, clickData.IpAddress,
		clickData.DeviceType, clickData.OperatingSystem, clickData.Referrer, clickData.Browser)
	return err
}
