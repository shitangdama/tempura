package database

import (
	"database/sql"
	"fmt"
	"os"

	// "database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB is
var DB *DBManager

// DBManager is
type DBManager struct {
	Conn *sql.DB
}

// DBInit is
func DBInit() {
	var err error
	conn, err := sql.Open("pgx", "postgres://postgres:kbr199sd5shi@localhost:5432/postgres")
	DB = &DBManager{
		Conn: conn,
	}
	// Conn, err = pgx.Connect(context.Background(), "postgres://postgres:kbr199sd5shi@localhost:5432/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err = DB.Conn.Ping(); err != nil {
		// do something about db error
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	fmt.Println("connect to database")

	_, err = DB.Conn.Exec("select 1;")
}

// DBClose is
func DBClose() {
	DB.Conn.Close()
}
