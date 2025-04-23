// cmd/main.go

package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
