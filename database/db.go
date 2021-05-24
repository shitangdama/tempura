package database

import (
	"context"
	"fmt"
	"gometa/models"
	"os"

	"github.com/jackc/pgx/v4"
	// "database/sql"
	// _ "github.com/jackc/pgx/v4/stdlib"
)

// Conn is
// var Conn *sql.DB
var Conn *pgx.Conn

// InitDB is
func InitDB() {
	var err error
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	Conn, err = pgx.Connect(context.Background(), "postgres://postgres:kbr199sd5shi@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}

// GetTables is
func GetTables() []models.Table {
	var tables []models.Table
	rows, err := Conn.Query(context.Background(), TabelSQL)
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
	return tables
}

func GetSchemas() {
	rows, err := Conn.Query(context.Background(), SchemaSql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	for rows.Next() {
		var id int32
		var name string
		var owner string
		err := rows.Scan(&id, &name, &owner)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		}
		fmt.Printf("%d. %s\n", id, name)
	}
}
