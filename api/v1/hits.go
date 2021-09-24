package v1

import (
	"context"
	"hits/api/prisma/db"
	"hits/api/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetHits(c *fiber.Ctx) error {
	var url = c.Params("url")
	var svg, _ = strconv.ParseBool(c.Query("svg"))
	var client = utils.GetPrisma()
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
			return c.Status(500).JSON(Response{
				Success: false,
				Message: "An internal server error occurred!",
			})
		}

		hit = createHit

	} else if err != nil {
		return c.Status(500).JSON(Response{
			Success: false,
			Message: "An internal server error occurred!",
		})
	}

	if svg == true {

		/*
			Put all the svg stuff in here
			We will want something like this https://github.com/ajstarks/svgo
			just go get github.com/ajstarks/svgo if you want to use this one
		*/

		return c.JSON(Response{
			Success: false,
			Message: "Mate",
		})
	}

	return c.JSON(Response{
		Success: true,
		Message: "Successfully requested hit!",
		Data:    hit,
	})
}
