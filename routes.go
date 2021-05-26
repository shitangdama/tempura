package main

// // Return all users as JSON
// func GetAllUsers(db *database.Database) fiber.Handler {
import (
	"gometa/app/controller"

	"github.com/gofiber/fiber/v2"
)

func registerApi(api fiber.Router) {
	roles := api.Group("/dbs")

	roles.Get("/tables", controller.GetAllTabels())
	// roles.Get("/:id", db.GetRole(db))

}
