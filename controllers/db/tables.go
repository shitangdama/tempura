package db

import (
	"gometa/database"

	"github.com/gofiber/fiber/v2"
)

// GetAllTabels all users as JSON
func GetAllTabels() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		result := database.GetTables()

		err := ctx.JSON(result)
		if err != nil {
			panic("Error occurred when returning JSON of a role: " + err.Error())
		}
		return err
	}
}
