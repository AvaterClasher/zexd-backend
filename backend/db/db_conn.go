package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error: No .env file found")
	}
}

var db *sql.DB

func CreateCon() *sql.DB {
	var (
		dbHost, _  = os.LookupEnv("DB_HOST")
		dbPort, _     = os.LookupEnv("DB_PORT")
		dbUser, _  = os.LookupEnv("DB_USER")
		dbPassword, _ = os.LookupEnv("DB_PASSWORD")
		dbName, _     = os.LookupEnv("DB_NAME")
	)
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open a connection: ", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping the database: ", err)
	}
	log.Println("Connected to PostgreSQL!")
	return db
}
