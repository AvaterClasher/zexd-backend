package daos

import (
	"database/sql"
	database "zexd/db"
	"zexd/logger"
	"zexd/model"

	_ "github.com/lib/pq"
)

var DB_NAME string
var TABLE_NAME string
var db *sql.DB

var log = logger.NewLogger()

func init() {
	db = database.CreateCon()

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS shortened_url (
			uid SERIAL PRIMARY KEY,
			user_id TEXT NOT NULL,
			url TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			clicks INT DEFAULT 0
		)
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}

func InsertShortenedUrl(urlModel model.UrlModel) (int, error) {
	var uid int
	query := "INSERT INTO shortened_url (url,user_id,created_at,clicks) VALUES ($1,$2,NOW(),0) RETURNING uid"
	err := db.QueryRow(query, urlModel.Url, urlModel.User_id).Scan(&uid)
	if err != nil {
		log.Errorf("Error inserting shortened URL: %v", err)
		return 0, err
	}
	return uid, nil
}
