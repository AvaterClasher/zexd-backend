package daos

import "log"

func IncrementClickCount(uid int) error {
	_, err := db.Exec("UPDATE shortened_url SET clicks = clicks + 1 WHERE uid = $1", uid)
	if err != nil {
		log.Println("Error updating click count:", err)
		return err
	}
	return nil
}
