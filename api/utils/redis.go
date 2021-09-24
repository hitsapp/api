package utils

import (
	"os"
	"github.com/gofiber/storage/redis"
)

func GetRedis() *redis.Storage {
	return redis.New(redis.Config{
		URL:   os.Getenv("REDIS_URI"),
		Reset: false,
	})
}
