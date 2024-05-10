package communities

import "github.com/gofiber/fiber"

func XService() (int, bool) {
	return fiber.StatusOK, true
}
