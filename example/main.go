package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jenggo/fiberlog"
)

func main() {
	app := fiber.New()

	// Default
	app.Use(fiberlog.New())

	app.Get("/ok", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Get("/warn", func(c *fiber.Ctx) error {
		return fiber.ErrUnprocessableEntity
	})

	app.Get("/err", func(c *fiber.Ctx) error {
		return fiber.ErrInternalServerError
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
