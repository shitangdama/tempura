package models

import (
	"fmt"
	"os"

	"gometa/database"
)

// Table is
type Table struct {
	ID               int32  `db:"id" json:"id"`
	Schema           string `db:"schema" json:"schema"`
	Name             string `db:"name" json:"name"`
	RlsEnabled       bool   `db:"rls_enabled" json:"rls_enabled"`
	RlsForced        bool   `db:"rls_forced" json:"rls_forced"`
	Relreplident     uint8  `db:"replica_identity" json:"replica_identity"`
	Bytes            int64  `db:"bytes" json:"bytes"`
	Size             int64  `db:"size" json:"size"`
	LiveRowsEstimate int    `db:"live_rows_estimate" json:"live_rows_estimate"`
	DeadRowsEstimate int    `db:"dead_rows_estimate" json:"dead_rows_estimate"`
	Comment          string `db:"comment" json:"comment"`
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
		// table := map[string]interface{}{}
		// if err = ScanMap(rows, &table); err != nil {
		if err = ScanStruct(rows, &table); err != nil {
			// tables = append(tables, table)
			fmt.Println(table)
		}
	}

	return tables
}
