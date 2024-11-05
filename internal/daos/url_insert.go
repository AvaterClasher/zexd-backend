package daos

import (
	"context"

	"github.com/AvaterClasher/zexd/internal/model"
	database "github.com/AvaterClasher/zexd/pkg/db"
	"github.com/AvaterClasher/zexd/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var DB_NAME string
var TABLE_NAME string
var db *pgxpool.Pool

var log = logger.NewLogger()

func init() {
	db = database.CreateCon()

	createUrlTableQuery := `
		CREATE TABLE IF NOT EXISTS shortened_url (
			uid SERIAL PRIMARY KEY,
			user_id TEXT NOT NULL,
			url TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT NOW(),
			clicks INT DEFAULT 0
		);
	`
	createClicksTableQuery := `
		CREATE TABLE IF NOT EXISTS click_metadata (
    		id SERIAL PRIMARY KEY,
    		url_uid INT REFERENCES shortened_url(uid) ON DELETE CASCADE,
    		clicked_at TIMESTAMP NOT NULL,
    		ip_address VARCHAR(45),
    		device_type VARCHAR(20),
    		operating_system VARCHAR(50),
    		referrer TEXT,
    		browser VARCHAR(50)
		);
	`
	_, err := db.Exec(context.Background(), createUrlTableQuery)
	if err != nil {
		log.Fatalf("Error creating URL table: %v", err)
	}
	_, err = db.Exec(context.Background(), createClicksTableQuery)
	if err != nil {
		log.Fatalf("Error creating Clicks table: %v", err)
	}
}

func InsertShortenedUrl(urlModel model.UrlModel) (int, error) {
	var uid int
	query := "INSERT INTO shortened_url (url,user_id,created_at,clicks) VALUES ($1,$2,NOW(),0) RETURNING uid"
	err := db.QueryRow(context.Background(), query, urlModel.Url, urlModel.User_id).Scan(&uid)
	if err != nil {
		log.Errorf("Error inserting shortened URL: %v", err)
		return 0, err
	}
	return uid, nil
}
