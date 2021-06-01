package controller

// // GetTables is
// func GetTables() []models.Table {
// 	var tables []models.Table
// 	rows, err := Conn.Query(context.Background(), TabelSQL)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 	}

// 	for rows.Next() {
// 		var id int32
// 		var schema string
// 		var name string
// 		var rls_enabled bool
// 		var rls_forced bool
// 		var bytes int32
// 		var size int32
// 		var live_rows_estimate int
// 		var dead_rows_estimate int
// 		var comment string

// 		err := rows.Scan(&id, &schema, &name, &rls_enabled, &rls_forced, &bytes, &size, &live_rows_estimate, &dead_rows_estimate, &comment)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		}
// 		fmt.Printf("%d. %s\n", id, schema)
// 	}
// 	return c.JSON(rows)
// }

// func createTable(c *fiber.Ctx) error {
// 	// https://github.com/datalanche/node-pg-format
// 	// 三个参数
// 	var name = c.Query("name", "")
// 	var schema = c.Query("name", "public")
// 	var comment = c.Query("name", "")
// 	// err := models.CreateTable(name, schema, comment)

// 	// retrieve
// 	return err
// }

// // TruncateTable is
// func TruncateTable(c *fiber.Ctx) error {
// 	// return DeleteTable(db, schema, tableName, true)

// 	return nil
// }

// // DropTable is
// func DropTable(c *fiber.Ctx) error {
// 	return nil
// }

// // AlterTable is
// func AlterTable(c *fiber.Ctx) error {
// 	return nil
// }
