package utils

import (
	"github.com/gofiber/storage/redis"
)

func GetRedis() *redis.Storage {
	store := redis.New()

	// Initialize custom config
	store = redis.New(redis.Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Username: "",
		Password: "",
		Database: 0,
		Reset:    false,
	})

	// or just the url with all information
	store = redis.New(redis.Config{
		URL:   "redis://localhost:6379",
		Reset: false,
	})

	return store
}
