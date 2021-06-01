package models

import (
	"database/sql"
	"fmt"
	"os"

	"gometa/database"
)

// Table is
type Table struct {
	ID               int32          `json:"id"`
	Schema           string         `json:"schema"`
	Name             string         `json:"name"`
	RlsEnabled       bool           `json:"rls_enabled"`
	RlsForced        bool           `json:"rls_forced"`
	Relreplident     uint8          `json:"replica_identity"`
	Bytes            int64          `json:"bytes"`
	Size             int64          `json:"size"`
	LiveRowsEstimate int            `json:"live_rows_estimate"`
	DeadRowsEstimate int            `json:"dead_rows_estimate"`
	Comment          sql.NullString `json:"comment"`
}

// CreateTable is
func CreateTable(DB *database.DBManager, name string, schema string, comment string) error {
	if schema == "" {
		schema = "public"
	}
	var tableSQL = fmt.Sprintf(`CREATE TABLE "%s"."%s" ();`, schema, name)

	var commentSQL = ""
	if comment != "" {
		commentSQL = fmt.Sprintf(`COMMENT ON TABLE "%s"."%s" IS "%s";`, schema, name, comment)
	}
	var SQL = fmt.Sprintf(`BEGIN; %s %s COMMIT;`, tableSQL, commentSQL)
	_, err := DB.Conn.Exec(SQL)

	// check iteration error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// // RetrieveTable is
// func RetrieveTable(name string, schema string) {

// }

// GetTables is
func GetTables(DB *database.DBManager) []Table {
	var tables []Table

	// invoke query
	rows, err := DB.Conn.Query(database.TabelSQL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	result := make(map[string]interface{})
	_ = ScanMap(rows, &result)
	// defer close result set
	defer rows.Close()

	fmt.Println(result)

	// 	// Iter results
	// for rows.Next() {

	// if err = rows.Scan(&n); err != nil {
	// 			fmt.Println(err) // Handle scan error
	// }
	// }

	// 	// check iteration error
	// 	if rows.Err() != nil {
	// 		fmt.Println(err)
	// 	}
	return tables
}

// https://blog.csdn.net/weixin_34168700/article/details/90433680

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
