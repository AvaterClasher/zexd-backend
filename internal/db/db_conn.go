package database

import (
	"context"
	"fmt"
	"os"
	"github.com/AvaterClasher/zexd/internal/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	log    = logger.NewLogger()
	db     *pgxpool.Pool
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

func CreateCon() *pgxpool.Pool {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.host, config.port, config.user, config.password, config.dbname, config.dbssl)

	var err error
	if db, err = pgxpool.New(context.Background(), connStr); err != nil {
		log.Fatalf("Failed to open a connection: %s", err)
	}
	if err = db.Ping(context.Background()); err != nil {
		log.Fatalf("Failed to ping the database: %s", err)
	}

	log.Info("Connected to PostgreSQL!")
	return db
}