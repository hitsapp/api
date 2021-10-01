package v1

import (
	"context"
	"fmt"
	. "hits/api/badge"
	"hits/api/prisma/db"
	"hits/api/utils"
	. "hits/api/utils"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func editText(text string) string {
	var output = fmt.Sprintf("#%s", strings.Trim(text, "\""))
	return output
}

func trimQuotes(text string) string {
	return strings.Trim(text, "\"")
}

func GetHits(c *fiber.Ctx) error {
	var urlQuery = c.Query("url")
	var json, _ = strconv.ParseBool(c.Query("json"))
	var colorQuery = c.Query("color")
	var leftBgColorQuery = c.Query("bgLeft")
	var rightBgColorQuery = c.Query("bgRight")
	var borderQuery = c.Query("border")
	var labelQuery = c.Query("label")
	var client = utils.GetPrisma()
	var ctx = context.Background()
	const regex = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
	var incrementValue = 1

	/* Strip Query paramters out of URL */
	u, err := url.Parse(urlQuery)

	match, _ := regexp.MatchString(regex, urlQuery)

	if urlQuery != "" && !match || err != nil {
		return c.Status(400).JSON(Response{
			Success: false,
			Message: "Invalid URL",
		})
	} else if urlQuery == "" {
		return c.Status(400).JSON(Response{
			Success: false,
			Message: "URL cannot be empty",
		})
	}

	urlQuery = u.Scheme + "://" + u.Host + u.Path

	if rightBgColorQuery == "" {
		rightBgColorQuery = "2f3136"
	} 
	
	if leftBgColorQuery == "" {
		leftBgColorQuery = "555"
	}

	ip, _ := GetRedis().Get(urlQuery + c.IP())

	// check if IP is less than 1 in length
	if len(ip) < 1 {
		GetRedis().Set(urlQuery+c.IP(), []byte(c.IP()), 30*time.Second)
	} else {
		incrementValue = 0
	}

	hit, err := client.Hits.FindUnique(
		db.Hits.URL.Equals(urlQuery),
	).Update(
		db.Hits.Hits.Increment(incrementValue),
	).Exec(ctx)

	if err != nil && err.Error() == "ErrNotFound" {
		createHit, createHitError := client.Hits.CreateOne(
			db.Hits.URL.Set(urlQuery),
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

	if json == true {
		return c.JSON(Response{
			Success: true,
			Message: "Successfully requested hit!",
			Data:    hit,
		})
	}

	svg := GenerateBadge(trimQuotes(labelQuery), strconv.Itoa(hit.Hits), editText(colorQuery), editText(leftBgColorQuery), editText(rightBgColorQuery), trimQuotes(borderQuery))
	c.Set(fiber.HeaderContentType, "image/svg+xml;charset=utf-8")
	c.Set(fiber.HeaderCacheControl, "max-age=0, s-maxage=0, must-revalidate, no-cache, no-store")
	return c.Send(svg)
}
