package utils

import (
	"os"
	"time"

	"github.com/gofiber/storage/redis"
)

var redisdb *redis.Storage

func SetRedis() {
	if redisdb == nil {
		println("REDIS URI", os.Getenv("REDIS_URI"))

		println("waiting 150ms")
		time.Sleep(150 * time.Millisecond)
		println("waited 150ms, trying to connect now")

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
