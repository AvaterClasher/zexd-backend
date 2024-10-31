package database

import (
	"database/sql"
	"fmt"
	"os"
	"zexd/logger"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var log = logger.NewLogger()

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Warn("Error: No .env file found")
	}
}

func CreateCon() *sql.DB {
	var (
		dbHost, _     = os.LookupEnv("DB_HOST")
		dbPort, _     = os.LookupEnv("DB_PORT")
		dbUser, _     = os.LookupEnv("DB_USER")
		dbPassword, _ = os.LookupEnv("DB_PASSWORD")
		dbName, _     = os.LookupEnv("DB_NAME")
	)
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=allow",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open a connection: %s", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %s", err)
	}
	log.Info("Connected to PostgreSQL!")
	return db
}
