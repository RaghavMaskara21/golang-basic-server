package communities

import (
	"github.com/gofiber/fiber/v2"
)

func X(ctx *fiber.Ctx) error {
	// jwtUserDetails := middlewares.GetJwtUser(ctx)
	// userId := utils.StringToInt32(jwtUserDetails.UserId.(string))
	// influencerId := utils.StringToInt32(strings.TrimSpace(ctx.Query("influencerId")))
	status, response := XService()
	return ctx.Status(status).JSON(response)
}
