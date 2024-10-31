package daos

import (
	"time"

	_ "github.com/lib/pq"
)

func DeleteExpiredUrls(expiryThreshold time.Time) error {
	query := "DELETE FROM shortened_url WHERE created_at < $1"
	result, err := db.Exec(query, expiryThreshold)
	if err != nil {
		log.Errorf("Error deleting expired URLs: %v", err)
		return err
	}

	deletedCount, err := result.RowsAffected()
	if err != nil {
		log.Errorf("Error retrieving affected rows: %v", err)
		return err
	}

	log.Infof("Deleted %d expired URL(s)\n", deletedCount)
	return nil
}
