package controller

import (
	"context"
	"fmt"
	"gometa/database"
	"gometa/models"
	"os"

	"github.com/gofiber/fiber/v2"
)

// GetTables is
func GetTables(c *fiber.Ctx) error {
	// var tables []models.Table
	rawRows, err := database.Conn.Query(context.Background(), database.TabelSQL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	var rows []models.Table
	for rawRows.Next() {
		var row models.Table
		// err := rawRows.(&row)
		ScanStruct(rawRows, &row)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		}
		rows = append(rows, row)
	}
	return c.JSON(rows)
}

func createTable(c *fiber.Ctx) error {
	// https://github.com/datalanche/node-pg-format
	// 三个参数
	var name = c.Query("name", "")
	var schema = c.Query("name", "public")
	var comment = c.Query("name", "")
	// err := models.CreateTable(name, schema, comment)

	// retrieve
	return err
}

// TruncateTable is
func TruncateTable(c *fiber.Ctx) error {
	// return DeleteTable(db, schema, tableName, true)

	return nil
}

// DropTable is
func DropTable(c *fiber.Ctx) error {
	return nil
}

// AlterTable is
func AlterTable(c *fiber.Ctx) error {
	return nil
}
