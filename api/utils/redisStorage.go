package utils

import (
	"os"

	"github.com/gofiber/storage/redis"
)

var redisdb *redis.Storage

func SetRedis() {
	if redisdb == nil {
		redisdb = redis.New(redis.Config{
			URL:   os.Getenv("REDIS_URI"),
			Reset: false,
		})
	}
}

func GetRedis() *redis.Storage {
	if redisdb == nil {
		SetRedis()
	}

	return redisdb
}
