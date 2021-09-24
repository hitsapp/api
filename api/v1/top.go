package v1

import (
	"context"
	"hits/api/prisma/db"
	. "hits/api/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetTopHits(c *fiber.Ctx) error {
	var reqTake = c.Query("take")
	var client = GetPrisma()
	var ctx = context.Background()

	/* If requester didn't put take in request */
	if len(reqTake) <= 0 {
		reqTake = "0"
	}

	take, err := strconv.Atoi(reqTake)
	if err != nil {
		return c.Status(400).JSON(Response{
			Success: false,
			Message: "Your request is improper!",
		})
	}

	/* Don't let requester go over 10 or under 0 */
	if take > 10 || take <= 0 {
		take = 5
	}

	hits, err := client.Hits.FindMany().OrderBy(
		db.Hits.Hits.Order(db.DESC),
	).Take(take).Exec(ctx)

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
