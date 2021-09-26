package v1

import (
	"context"
	"hits/api/prisma/db"
	. "hits/api/utils"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func GetTopHits(c *fiber.Ctx) error {
	var limit, _ = strconv.Atoi(c.Query("limit"))
	var client = GetPrisma()
	var ctx = context.Background()

	if limit == 0 {
		limit = 10
	}
	
	hits, err := client.Hits.FindMany().OrderBy(
		db.Hits.Hits.Order(db.DESC),
	).Take(limit).Exec(ctx)

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
