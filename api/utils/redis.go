package utils

import (
	"github.com/gofiber/storage/redis"
	"os"
)

func GetRedis() *redis.Storage {
	return redis.New(redis.Config{
		URL:   os.Getenv("REDIS_URI"),
		Reset: false,
	})
}
