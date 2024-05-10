package fiber

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func fiberErrorHandler(ctx *fiber.Ctx, err error) error {
	switch err.Error() {
	case "INTERNAL_SERVER_ERROR":
		return ctx.Status(fiber.StatusInternalServerError).JSON(INTERNAL_SERVER_ERROR)
	case "UNAUTHORIZED_ERROR":
		return ctx.Status(fiber.StatusUnauthorized).JSON(UNAUTHORIZED_ERROR)
	case "INVALID_GAME_ID_MSG":
		return ctx.Status(fiber.StatusBadRequest).JSON(INVALID_GAME_ID_MSG)
	case "BAD_REQUEST_ERROR":
		return ctx.Status(fiber.StatusBadRequest).JSON(BAD_REQUEST_ERROR)
	case "WRONG_CAROUSEL_TYPE":
		return ctx.Status(fiber.StatusBadRequest).JSON(WRONG_CAROUSEL_TYPE)
	case "CONFLICT_ERROR":
		return ctx.Status(fiber.StatusConflict).JSON(CONFLICT_ERROR)
	case "INFO_NOT_FOUND":
		return ctx.Status(fiber.StatusNotFound).JSON(INFO_NOT_FOUND)
	case "DECK_NOT_FOUND":
		return ctx.Status(fiber.StatusNotFound).JSON(DECK_NOT_FOUND)
	case "USER_CARD_NOT_FOUND":
		return ctx.Status(fiber.StatusNotFound).JSON(USER_CARD_NOT_FOUND)
	case "PROFILE_NOT_FOUND":
		return ctx.Status(fiber.StatusNotFound).JSON(PROFILE_NOT_FOUND)
	case "USER_NOT_FOUND":
		return ctx.Status(fiber.StatusNotFound).JSON(USER_NOT_FOUND)
	case "FLASH_REWARD_NOT_FOUND":
		return ctx.Status(fiber.StatusNotFound).JSON(FLASH_REWARD_NOT_FOUND)
	case "NO_SHOWCASE_DECK_FOUND":
		return ctx.Status(fiber.StatusNotFound).JSON(NO_SHOWCASE_DECK_FOUND)
	}

	//Default error code and message....
	code := fiber.StatusInternalServerError
	msg := fiber.ErrInternalServerError.Message

	//Override the error code and message if you found on the fiber err
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		msg = e.Message
	}
	return ctx.Status(code).SendString(msg)
}
