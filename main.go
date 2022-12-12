package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	//start a new fiber app
	app := fiber.New()

	//send a string back for GET calls to the endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})

	app.Listen(":3000")
}
