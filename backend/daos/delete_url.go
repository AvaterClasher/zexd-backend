package daos

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DeleteUrl(uid int) error {
	query := "DELETE FROM shortened_url WHERE uid = $1"
	result, err := db.Exec(query, uid)
	if err != nil {
		log.Println("Error deleting from database:", err)
		return err
	}

	deletedCount, err := result.RowsAffected()
	if err != nil {
		log.Println("Error retrieving affected rows:", err)
		return err
	}

	fmt.Printf("Deleted %v url(s)\n", deletedCount)
	return nil
}
