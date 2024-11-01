package cache

import (
	"context"
	"os"
	"zexd/logger"

	"github.com/joho/godotenv"
	redis "github.com/redis/go-redis/v9"
)

var (
	log    = logger.NewLogger()
	rdb    *redis.Client
	Ctx    = context.Background()
	domain string
)

func init() {
	loadRedisDomain()
}

func loadRedisDomain() {
	if domain = os.Getenv("REDIS_DOMAIN"); domain == "" {
		if err := godotenv.Load(".env"); err != nil {
			log.Warn("Error: No .env file found")
		}
		domain = os.Getenv("REDIS_DOMAIN")
	}

	if domain == "" {
		log.Warn("Warning: REDIS_DOMAIN environment variable is not set")
	}
}

func CreateCon() *redis.Client {
	if rdb != nil {
		return rdb
	}

	opt, err := redis.ParseURL(domain)
	if err != nil {
		log.Fatalf("Failed to parse Redis URL: %v", err)
	}

	rdb = redis.NewClient(opt)

	err = rdb.Set(Ctx, "http://localhost/Avatar", "https://soumyadipmoni.vercel.app", 0).Err()
	if err != nil {
		log.Fatalf("Failed to set value in Redis: %v", err)
	}

	log.Info("Connected to Redis Container!")
	return rdb
}
