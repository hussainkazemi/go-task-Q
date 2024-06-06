package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupAPIV1(app *fiber.App) {
	router := app.Group("/api/v1")

	router.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}
