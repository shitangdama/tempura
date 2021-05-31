package database

import (
	"database/sql"
	"fmt"
	"os"

	// "database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Conn is
// var Conn *sql.DB
var Conn *sql.DB

// InitDB is
func InitDB() {
	var err error
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	// Go database/sql 教程_weixin_34168700的博客-CSDN博客
	Conn, err := sql.Open("pgx", "postgres://postgres:kbr199sd5shi@localhost:5432/postgres")

	// Conn, err = pgx.Connect(context.Background(), "postgres://postgres:kbr199sd5shi@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err = Conn.Ping(); err != nil {
		// do something about db error
	}
}
