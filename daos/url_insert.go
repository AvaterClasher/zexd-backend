package daos

import (
	"context"
	database "zexd/db"
	"zexd/logger"
	"zexd/model"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var DB_NAME string
var TABLE_NAME string
var db *pgxpool.Pool

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

	_, err := db.Exec(context.Background(), createTableQuery)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}

func InsertShortenedUrl(urlModel model.UrlModel) (int, error) {
	var uid int
	query := "INSERT INTO shortened_url (url,user_id,created_at,clicks) VALUES ($1,$2,NOW(),0) RETURNING uid"
	err := db.QueryRow(context.Background(),query, urlModel.Url, urlModel.User_id).Scan(&uid)
	if err != nil {
		log.Errorf("Error inserting shortened URL: %v", err)
		return 0, err
	}
	return uid, nil
}
