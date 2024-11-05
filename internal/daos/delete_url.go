package daos

import (
	"context"

	_ "github.com/lib/pq"
)

func DeleteUrl(uid int) error {
	query := "DELETE FROM shortened_url WHERE uid = $1"
	result, err := db.Exec(context.Background(),query, uid)
	if err != nil {
		log.Errorf("Error deleting URL: %v", err)
		return err
	}

	deletedCount := result.RowsAffected()

	log.Infof("Deleted %v url(s)\n", deletedCount)
	return nil
}