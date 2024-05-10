package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:       "main-service-v2",
		StrictRouting: true,
		ErrorHandler:  fiberErrorHandler,
	})

	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${time} ${ip} ${reqHeader:X-Forwarded-For} ${locals:requestid} ${status} - ${method} ${latency} ${path} ${body}\n",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	//Add the accept-encoding header to gzip for data compression
	app.Use(func(c *fiber.Ctx) error {
		c.Request().Header.Add("Accept-Encoding", "gzip")
		return c.Next()
	})

	app.Use(pprof.New(pprof.Config{Prefix: "/api/v4/monitoring/performance"}))

	return app
}
