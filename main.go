package main

import (
	//"hayday/server/config"
	"hayday/server/config"
	"hayday/server/internal/server"
	"hayday/server/internal/websocket"
	"hayday/server/src/router"
	"strings"
)

func main() {
	envValues := config.LoadConfig()
	if strings.ToUpper(envValues.APP_ENV) != "PROD" {
		config.FIRESTORE_CLUB_COLL = "STAGE_CLUB_ROOMS"
		config.FIRESTORE_CREATOR_WAITLISTING_COLL = "STAGE_CREATOR_WAITLISTING"
	}
	server.InitiateConnections()
	router.SetupRouter()
	server.StartServer(config.EnvValues.SERVER_PORT)
	websocket.InitiateSocketServer()
}
