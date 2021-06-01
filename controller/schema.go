package controller

// GetSchemas is
// func GetSchemas(c *fiber.Ctx) error {
// 	// var tables []models.Table
// 	rawRows, err := database.Conn.Query(context.Background(), database.SchemaSQL)
// 	var rows []models.Schema
// 	for rawRows.Next() {
// 		var row models.Schema
// 		// err := rawRows.(&row)
// 		ScanStruct(rawRows, &row)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		}
// 		rows = append(rows, row)
// 	}
// 	return c.JSON(rows)
// }
