package v1

import (
	"context"
	"hits/api/prisma/db"
	. "hits/api/utils"

	"github.com/gofiber/fiber/v2"
)

func Introduction(c *fiber.Ctx) error {
	return c.JSON(Response{
		Success: true,
		Message: "Welcome to hits API V1!",
	})
}

func GetHits(c *fiber.Ctx) error {
	var url = c.Params("url")
	var client = GetPrisma()
	var ctx = context.Background()

	hit, err := client.Hits.FindUnique(
		db.Hits.URL.Equals(url),
	).Update(
		db.Hits.Hits.Increment(1),
	).Exec(ctx)

	if err != nil && err.Error() == "ErrNotFound" {
		createHit, createHitError := client.Hits.CreateOne(
			db.Hits.URL.Set(url),
		).Exec(ctx)

		if createHitError != nil {
			return c.JSON(Response{
				Success: false,
				Message: "An internal server error occurred!",
			})
		}

		hit = createHit

	} else if err != nil {
		return c.JSON(Response{
			Success: false,
			Message: "An internal server error occurred!",
		})
	}

	return c.JSON(Response{
		Success: true,
		Message: url,
		Data:    hit,
	})
}

func GetTopHits(c *fiber.Ctx) error {
	var take = c.Query("take")
	println(take)

	return c.JSON(Response{
		Success: true,
		Message: "Here are the things",
	})
}
