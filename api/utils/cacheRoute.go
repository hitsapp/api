package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func CacheRoute() fiber.Handler {
	return cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1" || c.IP() == "localhost"
		},
		Expiration:   30 * time.Second,
		CacheControl: true,
		Storage:      GetRedis(),
	})
}
