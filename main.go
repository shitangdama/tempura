package main

import (
	"fmt"

	"gometa/database"
	"gometa/models"
	// _ "github.com/jackc/pgx/v4/stdlib"
)

// // SetupRoutes is
// func SetupRoutes(app *fiber.App) {
// 	dbRoutes := app.Group("/db")

// 	dbRoutes.Get("/tables", controller.GetTables)
// 	dbRoutes.Get("/columns", controller.GetColumns)
// 	// roles.Get("/:id", db.GetRole(db))

// }

func main() {
	database.DBInit()
	// app := fiber.New()
	// app.Use(cors.New())

	// database.GetSchemas()
	// SetupRoutes(app)

	// app.Listen(":3000")
	fmt.Print(database.DB)
	// models.CreateTable(database.DB, "test", "", "")
	models.GetTables(database.DB)

}

// we need some data to query; I'll create a temp table
// and populate
