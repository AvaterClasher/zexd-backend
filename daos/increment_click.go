package daos

func IncrementClickCount(uid int) error {
	_, err := db.Exec("UPDATE shortened_url SET clicks = clicks + 1 WHERE uid = $1", uid)
	if err != nil {
		log.Errorf("Error updating click count: %v", err)
		return err
	}
	return nil
}
