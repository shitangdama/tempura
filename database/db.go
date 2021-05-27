package database

import (
	"context"
	"fmt"
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
