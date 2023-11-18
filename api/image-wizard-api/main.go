package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋! from API")
	})

	app.Listen(":5000")
}
