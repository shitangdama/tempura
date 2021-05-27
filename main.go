package main

import (
	"gometa/controller"
	"gometa/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// SetupRoutes is
func SetupRoutes(app *fiber.App) {
	dbRoutes := app.Group("/db")

	dbRoutes.Get("/tables", controller.GetTables)
	// roles.Get("/:id", db.GetRole(db))

}

func main() {
	database.InitDB()
	app := fiber.New()
	app.Use(cors.New())

	// database.GetSchemas()

	SetupRoutes(app)

	app.Listen(":3000")
}
