package v1

import (
	netURL "net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hitsapp/api/utils"
)

const urlRegex = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`

func Hits(c *fiber.Ctx) error {
	var url = c.Query("url")
	var color = c.Query("color")
	var leftBgColor = c.Query("bgLeft")
	var rightBgColor = c.Query("bgRight")
	var border = c.Query("border")
	var label = c.Query("label")

	var wantsJSON = string(c.Request().Header.Peek("Accept")) == "application/json"

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	parsedURL, err := netURL.Parse(url)
	url = parsedURL.Scheme + "://" + parsedURL.Host + parsedURL.Path
	matched, matchErr := regexp.MatchString(urlRegex, url)

	if err != nil || url == "" || matchErr != nil || !matched {
		return c.Status(400).JSON(utils.Response[any]{
			Success: false,
			Error: &utils.Error{
				Code:    "invalid_url",
				Message: "You have passed in an invalid URL.",
			},
		})
	}

	if color == "" {
		color = "fff"
	}

	if leftBgColor == "" {
		leftBgColor = "555"
	}

	if rightBgColor == "" {
		rightBgColor = "2f3136"
	}

	if border == "" {
		border = "000000"
	}

	if label == "" {
		label = "hits"
	}

	ip := utils.GetKey(c.IP())
	if ip == "" {
		utils.IncrementHit(url)
		utils.AddKey(c.IP(), "1", 30)
	}

	urlHits := utils.GetHit(url)
	if urlHits == 0 {
		utils.AddHit(url)
		urlHits = 1
	}

	if wantsJSON {
		return c.JSON(utils.Response[int]{
			Success: true,
			Data:    int(urlHits),
		})
	}

	svg := utils.GenerateBadge(utils.TrimQuotes(label), strconv.Itoa(int(urlHits)), utils.TransformToHex(color), utils.TransformToHex(leftBgColor), utils.TransformToHex(rightBgColor), utils.TrimQuotes(border))
	c.Set(fiber.HeaderContentType, "image/svg+xml")
	c.Set(fiber.HeaderCacheControl, "max-age=0, s-maxage=0, must-revalidate, no-cache, no-store")

	return c.SendString(string(svg))
}
