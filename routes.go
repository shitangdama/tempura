package main

import (
	"gometa/controller"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes is
func SetupRoutes(app *fiber.App) {
	dbRoutes := app.Group("/db")

	dbRoutes.Get("/tables", controller.GetTables)
	// roles.Get("/:id", db.GetRole(db))

}
