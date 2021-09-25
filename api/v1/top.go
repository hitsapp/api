package v1

import (
	"context"
	"hits/api/prisma/db"
	. "hits/api/utils"

	"github.com/gofiber/fiber/v2"
)

func GetTopHits(c *fiber.Ctx) error {
	var client = GetPrisma()
	var ctx = context.Background()

	hits, err := client.Hits.FindMany().OrderBy(
		db.Hits.Hits.Order(db.DESC),
	).Take(10).Exec(ctx)

	if err != nil {
		return c.Status(500).JSON(Response{
			Success: false,
			Message: "An internal server error occurred whilst fullfilling your request!",
		})
	}

	return c.JSON(Response{
		Success: true,
		Message: "Successfully retrieved hits leaderboard",
		Data:    hits,
	})
}
