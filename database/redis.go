package database

import (
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func StartRedis() *redis.Client {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := os.Getenv("REDIS_DB")

	opt := &redis.Options{
		Addr:     redisURL,
		Password: redisPassword,
		DB:       0,
	}

	if redisDB != "" {
		db, err := strconv.Atoi(redisDB)
		if err != nil {
			panic(err)
		}
		opt.DB = db
	}

	redisClient := redis.NewClient(opt)
	return redisClient
}
