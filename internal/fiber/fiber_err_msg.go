package fiber

import "github.com/gofiber/fiber/v2"

var (
	BAD_REQUEST_ERROR = fiber.Map{
		"statusCode": 400,
		"message":    "Something went wrong!",
		"error":      "Bad Request",
	}

	INVALID_GAME_ID_MSG = fiber.Map{
		"statusCode": 400,
		"message":    "Invalid game id!",
		"error":      "Bad Request",
	}

	INTERNAL_SERVER_ERROR = fiber.Map{
		"statusCode": 500,
		"message":    "Something went wrong!",
		"error":      "Something went wrong!",
	}

	UNAUTHORIZED_ERROR = fiber.Map{
		"statusCode": 401,
		"message":    "Unauthorized",
		"error":      "Unauthorized",
	}

	WRONG_CAROUSEL_TYPE = fiber.Map{
		"statusCode": 400,
		"message":    "Invalid carousel type!",
		"error":      "Bad Request",
	}

	CONFLICT_ERROR = fiber.Map{
		"statusCode": 409,
		"message":    "Something went wrong! Internal error",
		"error":      "Conflict",
	}

	INFO_NOT_FOUND = fiber.Map{
		"statusCode": 404,
		"message":    "Info not found",
		"error":      "Not found",
	}

	DECK_NOT_FOUND = fiber.Map{
		"statusCode": 404,
		"message":    "Deck not found",
		"error":      "Not found",
	}

	USER_CARD_NOT_FOUND = fiber.Map{
		"statusCode": 404,
		"message":    "Card not found",
		"error":      "Not found",
	}

	PROFILE_NOT_FOUND = fiber.Map{
		"statusCode": 404,
		"message":    "Profile not found",
		"error":      "Not found",
	}

	USER_NOT_FOUND = fiber.Map{
		"statusCode": 404,
		"message":    "User not found",
		"error":      "Not found",
	}

	FLASH_REWARD_NOT_FOUND = fiber.Map{
		"statusCode": 404,
		"message":    "No flash reward available for today!",
		"error":      "Not found",
	}

	NO_SHOWCASE_DECK_FOUND = fiber.Map{
		"statusCode": 404,
		"message":    "Showcase deck not found",
		"error":      "Not found",
	}
)
