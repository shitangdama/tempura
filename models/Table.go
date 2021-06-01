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

	defer rows.Close()

	for rows.Next() {
		table := Table{}
		if err = ScanStruct(rows, &table); err != nil {
			tables = append(tables, table)
			fmt.Println(table.ID)
		}
	}

	return tables
}
