package server

import (
	"fmt"
	fiber2 "hayday/server/internal/fiber"
	"hayday/server/internal/mongodb"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

var server *Server

type Server struct {
	MongoClient *mongodb.MongoDbStorage
	FiberApp    *fiber.App
}

func Instance() *Server {
	return server
}

func InitiateConnections() *Server {

	mongoClient := mongodb.Connect()
	fiberApp := fiber2.InitFiber()
	fiberApp.Config()
	server = &Server{
		MongoClient: mongodb.NewMongoStorage(mongoClient),
		FiberApp:    fiberApp,
	}
	return server
}
func StartServer(serverPort string) {
	port := fmt.Sprintf(":%s", serverPort)
	err := server.FiberApp.Listen(port)
	if err != nil {
		fmt.Printf("failed to start the rest api server : ERROR : %s \n", err)
		err := syscall.Kill(syscall.Getpid(), syscall.SIGUSR1)
		if err != nil {
			fmt.Printf("failed to do syscall KILL : ERROR : %s \n", err)
			panic(err)
		}
	}
}
