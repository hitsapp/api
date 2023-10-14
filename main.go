package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hitsapp/api/utils"
	v1 "github.com/hitsapp/api/v1"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: false,
		StrictRouting: true,
		ServerHeader:  "hits",
		BodyLimit:     1024 * 1024,
		GETOnly:       true,
		ProxyHeader:   os.Getenv("PROXY_HEADER"),
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	utils.ConnectRedis()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "sponsored by https://hop.io",
		})
	})

	v1Group := app.Group("/v1")
	v1Group.Get("/leaderboard", utils.RateLimit(50), v1.Leaderboard)
	v1Group.Get("/hits", utils.RateLimit(1200), v1.Hits)

	log.Fatal(app.Listen(":8080"))
}
