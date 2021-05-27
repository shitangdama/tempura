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
	var tables []models.Table
	rows, err := database.Conn.Query(context.Background(), database.TabelSQL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	for rows.Next() {
		var id int32
		var schema string
		var name string
		var rls_enabled bool
		var rls_forced bool
		var bytes int32
		var size int32
		var live_rows_estimate int
		var dead_rows_estimate int
		var comment string

		err := rows.Scan(&id, &schema, &name, &rls_enabled, &rls_forced, &bytes, &size, &live_rows_estimate, &dead_rows_estimate, &comment)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		}
		fmt.Printf("%d. %s\n", id, schema)
	}
	return c.JSON(tables)
}
