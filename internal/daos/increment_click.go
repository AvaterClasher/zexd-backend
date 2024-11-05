package daos

import "context"

func IncrementClickCount(uid int) error {
	query := "UPDATE shortened_url SET clicks = clicks + 1 WHERE uid = $1"
	_, err := db.Exec(context.Background(), query, uid)
	if err != nil {
		log.Errorf("Error updating click count: %v", err)
		return err
	}
	return nil
}
