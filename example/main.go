package main

import (
	"github.com/dre1080/fiberlog"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Default
	app.Use(fiberlog.New())

	app.Get("/ok", func(c *fiber.Ctx) error {
		c.SendString("ok")
		return nil
	})

	app.Get("/warn", func(c *fiber.Ctx) error {
		return fiber.ErrUnprocessableEntity
	})

	app.Get("/err", func(c *fiber.Ctx) error {
		return fiber.ErrInternalServerError
	})

	app.Listen(":3000")
}
