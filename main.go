package main

import (
	"gometa/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// 先获取
// https://github.com/thomasvvugt/fiber-boilerplate/blob/48cf7a4d10d3a58ea3618a7ef46d8b16fcf2f1e9/main.go

func main() {
	database.InitDB()
	app := fiber.New()
	app.Use(cors.New())

	// database.GetSchemas()

	SetupRoutes(app)

	app.Listen(":3000")
}
