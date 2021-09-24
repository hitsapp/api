package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimit(amount int) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        amount,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(Response{
				Success: false,
				Message: "Ratelimit exceeded. Try again in 1 minute.",
			})
		},
		Storage: GetRedis(),
	})
}
