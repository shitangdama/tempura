package models

import (
	"context"
	"database/sql"
	"fmt"
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
func CreateTable(name string, schema string, comment string) error {
	var defaultSchema string
	if schema != "" {
		defaultSchema = schema
	}
	var tableSQL = fmt.Sprintf(`CREATE TABLE "%s"."%s"`, defaultSchema, name)
	if comment != "" {
		const commentSql = fmt.Sprintf(`CREATE TABLE "%s"."%s"`, defaultSchema, name)
	}
	_, err := database.Conn.Query(context.Background(), tableSQL)

	return err
}

func GetTables() {

	var sum, n int32

	// invoke query
	rows, err := db.Query("SELECT generate_series(1,$1)", 10)
	// handle query error
	if err != nil {
		fmt.Println(err)
	}
	// defer close result set
	defer rows.Close()

	// Iter results
	for rows.Next() {
		if err = rows.Scan(&n); err != nil {
			fmt.Println(err) // Handle scan error
		}
		sum += n // Use result
	}

	// check iteration error
	if rows.Err() != nil {
		fmt.Println(err)
	}

}

func UpdateTable() {

}

// https://blog.csdn.net/weixin_34168700/article/details/90433680
