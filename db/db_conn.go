package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	"zexd/logger"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	log = logger.NewLogger()
	db  *sql.DB
	config = struct {
		host, port, user, password, dbname, dbssl string
	}{}
)

var envVars = []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_SSL"}

func init() {
	loadEnv()
	initializeConfig()
}

func loadEnv() {
	for _, env := range envVars {
		if _, exists := os.LookupEnv(env); !exists {
			if err := godotenv.Load(".env"); err != nil {
				log.Warn("Error: No .env file found")
			}
			break
		}
	}
}

func initializeConfig() {
	config.host = os.Getenv("DB_HOST")
	config.port = os.Getenv("DB_PORT")
	config.user = os.Getenv("DB_USER")
	config.password = os.Getenv("DB_PASSWORD")
	config.dbname = os.Getenv("DB_NAME")
	config.dbssl = os.Getenv("DB_SSL")

	for i, value := range []string{config.host, config.port, config.user, config.password, config.dbname, config.dbssl} {
		if value == "" {
			log.Warn(fmt.Sprintf("Warning: Required environment variable %s is not set", envVars[i]))
		}
	}
}

func CreateCon() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.host, config.port, config.user, config.password, config.dbname, config.dbssl)

	var err error
	if db, err = sql.Open("postgres", connStr); err != nil {
		log.Fatalf("Failed to open a connection: %s", err)
	}

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(30)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %s", err)
	}

	log.Info("Connected to PostgreSQL!")
	return db
}