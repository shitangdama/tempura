package main

import (
	"gometa/pkg/database"

	"github.com/gofiber/fiber/v2"
	// "gometa/pkg/database"
)

// 先获取
// https://github.com/thomasvvugt/fiber-boilerplate/blob/48cf7a4d10d3a58ea3618a7ef46d8b16fcf2f1e9/main.go

func main() {
	app := fiber.New()
	database.InitDB()
	database.GetSchemas()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/tables", func(c *fiber.Ctx) error {
		return c.SendString("tables")
	})

	// app.Get("/schemas", func(c *fiber.Ctx) error {
	// 	return c.SendString("schemas")
	// })

	// app.Listen(":3000")
}
