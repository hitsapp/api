package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimit(amount int) fiber.Handler {
	return limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1" || c.IP() == "localhost"
		},
		Max:        amount,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(Response[*any]{
				Success: false,
				Error: &Error{
					Code:    "rate_limit_reached",
					Message: "You have reached the rate limit, please try again in one minute",
				},
			})
		},
	})
}
