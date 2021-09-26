package v1

import (
	"context"
	"github.com/ajstarks/svgo"
	"github.com/gofiber/fiber/v2"
	"hits/api/prisma/db"
	"hits/api/utils"
	. "hits/api/utils"
	"regexp"
	"strconv"
)

func GetHits(c *fiber.Ctx) error {
	var url = c.Query("url")
	var svgQuery, _ = strconv.ParseBool(c.Query("svg"))
	var client = utils.GetPrisma()
	var ctx = context.Background()
	const regex = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
	match, _ := regexp.MatchString(regex, url)

	if !match {
		return c.Status(400).JSON(Response{
			Success: false,
			Message: "Invalid URL",
		})
	}

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

	if svgQuery == true {
		c.Set("Content-Type", "image/svg+xml")
		s := svg.New(c)
		s.Start(500, 500)
		s.Square(250, 250, 125, "fill:none;stroke:black")
		s.End()
		return nil
	}

	return c.JSON(Response{
		Success: true,
		Message: "Successfully requested hit!",
		Data:    hit,
	})
}
