package main

import (
	"hits/api/prisma/db"
	. "hits/api/utils"
	. "hits/api/v1"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
		ServerHeader:  "HITS API",
		AppName:       "HITS API v1.0",
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
