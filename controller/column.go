package controller

// // GetColumns is
// func GetColumns(c *fiber.Ctx) error {
// 	// var tables []models.Table
// 	rawRows, err := database.Conn.Query(context.Background(), database.ColumnSQL)
// 	var rows []models.Column
// 	for rawRows.Next() {
// 		var row models.Column
// 		// err := rawRows.(&row)
// 		ScanStruct(rawRows, &row)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		}
// 		rows = append(rows, row)
// 	}
// 	return c.JSON(rows)
// }
