package main

// // Return all users as JSON
// func GetAllUsers(db *database.Database) fiber.Handler {
import (
	"gometa/app/controllers/db"

	"github.com/gofiber/fiber/v2"
)

func registerDBs(api fiber.Router) {
	roles := api.Group("/dbs")

	roles.Get("/", db.GetAllTabels())
	// roles.Get("/:id", db.GetRole(db))

}
