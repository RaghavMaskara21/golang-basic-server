package router

import (
	"hayday/server/internal/server"

	"hayday/server/src/api"
	"hayday/server/src/api/communities"
)

func SetupRouter() {
	app := server.Instance().FiberApp
	app.Get("/api/v4/health", api.HealthCheck)

	// Below Routers with Authentication....
	//router := app.Group("/api/v4", middlewares.JWTAuthValidator)
	router := app.Group("/api/v4")
	router.Get("/hero", communities.X)

}
