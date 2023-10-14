package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hitsapp/api/utils"
)

func Leaderboard(c *fiber.Ctx) error {
	leaderboard := utils.GetHits(20)

	var parsedLeaderboard []utils.Leaderboard
	for _, v := range leaderboard {
		parsedLeaderboard = append(parsedLeaderboard, utils.Leaderboard{
			URL:  v.Member.(string),
			Hits: int(v.Score),
		})
	}

	return c.JSON(utils.Response[[]utils.Leaderboard]{
		Success: true,
		Data:    parsedLeaderboard,
	})
}
