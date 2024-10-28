package daos

import (
	"log"
	"time"

	_ "github.com/lib/pq"
)

func DeleteExpiredUrls(expiryThreshold time.Time) error {
	query := "DELETE FROM shortened_url WHERE created_at < $1"
	result, err := db.Exec(query, expiryThreshold)
	if err != nil {
		log.Println("Error deleting expired URLs:", err)
		return err
	}

	deletedCount, err := result.RowsAffected()
	if err != nil {
		log.Println("Error retrieving affected rows:", err)
		return err
	}

	log.Printf("Deleted %d expired URL(s)\n", deletedCount)
	return nil
}
