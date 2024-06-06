package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-async-task-queue/api"
	"log"
)

func StartHttpServer() {
	app := fiber.New(fiber.Config{})
	api.SetupAPIV1(app)
	if err := app.Listen(fmt.Sprintf("%s:%s", "127.0.0.1", "2100")); err != nil {
		log.Fatalf("Could not start Fiber server: %v", err)
	}
}
