package models

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4"
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
func (c *pgx.Conn) CreateTable(name string, schema string, comment string) error {
	// 	var defaultSchema string
	// 	if schema != "" {
	// 		defaultSchema = schema
	// 	}
	// 	var tableSQL = fmt.Sprintf(`CREATE TABLE "%s"."%s"`, defaultSchema, name)
	// 	if comment != "" {
	// 		const commentSql = fmt.Sprintf(`CREATE TABLE "%s"."%s"`, defaultSchema, name)
	// 	}
	// 	_, err := database.Conn.Query(context.Background(), tableSQL)
	var tableSQL = fmt.Sprintf(`CREATE TABLE "%s"."%s" ();`, schema, name)
	var commentSQL = ""
	if comment != "" {
		commentSQL = fmt.Sprintf(`COMMENT ON TABLE "%s"."%s" IS "%s";`, schema, name, comment)
	}
	var SQL = fmt.Sprintf(`BEGIN; %s %s COMMIT;`, tableSQL, commentSQL)

	_, err := c.Exec(context.Background(), SQL)
	return err
}

// RetrieveTable is
func RetrieveTable(name string, schema string) {

}

// UpdateTable is
func UpdateTable() {

}
