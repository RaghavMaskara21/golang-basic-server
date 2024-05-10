package api

import (
	"github.com/gofiber/fiber/v2"
)

func HealthCheck(ctx *fiber.Ctx) error {
	// s := server.Instance()
	// isDbHealthy := mysql.IsDBHealthy(s.DbStorage.DB)
	// isRedisHealthy := redis.IsRedisHealthy(s.RedisClient)

	// if isRedisHealthy && isDbHealthy {
	return ctx.Status(fiber.StatusOK).SendString("Healthy")
	// }
	// return ctx.Status(fiber.StatusInternalServerError).SendString("Db or Redis is not healthy!")
}
