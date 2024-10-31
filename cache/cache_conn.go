package cache

import (
	"context"
	"os"
	"zexd/logger"

	"github.com/joho/godotenv"
	redis "github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var log = logger.NewLogger()

func init() {
	envFile := ".env"
	if err := godotenv.Load(envFile); err != nil {
		log.Warnf("Error: No Environment File Found, cache_connection.go %v", err)
	}
}

func tryRedis(domain string) *redis.Client {
	opt, _ := redis.ParseURL(domain)
	rdb := redis.NewClient(opt)

	inputUrl := "https://soumyadipmoni.vercel.app"
	newUrl := "http://localhost/Avater"

	err := rdb.Set(Ctx, newUrl, inputUrl, 0).Err()

	if err != nil {
		log.Error(err, "Current Domain", domain)
		return nil
	}

	return rdb
}

func CreateCon() *redis.Client {
	var cacheDomain = os.Getenv("REDIS_DOMAIN")

	client := tryRedis(cacheDomain)

	if client == nil {
		log.Error("Connection Failed while trying to connect with Redis.")
	} else {
		log.Info("Connected to Redis Container!!")
	}

	return client
}
