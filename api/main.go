package main

import (
	"hits/api/prisma/db"
	"hits/api/utils"
	. "hits/api/utils"
	. "hits/api/v1"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func Main(c *fiber.Ctx) error {
	return c.JSON(Response{
		Success: true,
		Message: "Welcome to hits API V1!",
	})
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(Response{
			Success: true,
			Message: "Welcome to Hits API!",
		})
	})

	/* V1 */
	v1 := app.Group("/v1")
	v1.Get("/", Main)
	v1.Get("/top", GetTopHits)
	v1.Get("/hits/:url", GetHits)
}

func main() {
	godotenv.Load()

	app := fiber.New(fiber.Config{
		CaseSensitive: false,
		StrictRouting: true,
		ServerHeader:  "Hits API",
		AppName:       "Hits API v1.0",
		BodyLimit:     1024 * 1024,
	})

	app.Use(logger.New(logger.Config{
		Format: "${time} |   ${cyan}${status} ${reset}|   ${latency} | ${ip} on ${cyan}${ua} ${reset}| ${cyan}${method} ${reset}${path} \n",
	}))

	app.Use(recover.New(recover.Config{
		Next:             nil,
		EnableStackTrace: true,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Use(limiter.New(limiter.Config{
		Max:        15,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.JSON(Response{
				Success: false,
				Message: "Ratelimit exceeded. Try again in 1 minute.",
			})
		},
		Storage: utils.GetRedis(),
	}))

	setupRoutes(app)

	if GetPrisma() == nil {
		SetGlobalDb(db.NewClient())
	}

	if err := GetPrisma().Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := GetPrisma().Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
