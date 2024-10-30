package daos

func UrlExists(uid int) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM shortened_url WHERE uid = $1)"
	err := db.QueryRow(query, uid).Scan(&exists)
	if err != nil {
		log.Errorf("Error checking URL existence in database: %v", err)
		return false, err
	}
	return exists, nil
}