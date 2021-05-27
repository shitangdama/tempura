package models

import "database/sql"

// Column is
type Column struct {
	ID                 int64          `json:"id"`
	TableID            int64          `json:"table_id"`
	Schema             string         `json:"schema"`
	Table              string         `json:"table"`
	OrdinalPosition    int64          `json:"ordinal_position"`
	Name               string         `json:"name"`
	ColumnDefault      sql.NullString `json:"default_value"`
	DataType           string         `json:"data_type"`
	Format             string         `json:"format"`
	IsIdentity         bool           `is_identityjson:"is_identity"`
	IdentityGeneration string         `json:"identity_generation"`
	IsNullable         bool           `json:"is_nullable"`
	IsUpdatable        bool           `json:"is_updatable"`
	Enums              []string       `json:"enums"`
	Comment            string         `json:"comment"`
}
