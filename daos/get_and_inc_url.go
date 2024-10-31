package daos

import (
	_ "github.com/lib/pq"
)

func GetUrlAndIncrement(uid int) (string, error) {
	var url string
	query := "SELECT url FROM shortened_url WHERE uid = $1"
	err := db.QueryRow(query, uid).Scan(&url)
	if err != nil {
		log.Errorf("Error retrieving URL: %v", err)
		return "", err
	}

	_, err = db.Exec("UPDATE shortened_url SET clicks = clicks + 1 WHERE uid = $1", uid)
	if err != nil {
		log.Errorf("Error updating click count: %v", err)
	}

	return url, nil
}