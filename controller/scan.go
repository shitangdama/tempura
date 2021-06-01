package controller

import (
	"github.com/jackc/pgx/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// FieldNames is
func FieldNames(rows pgx.Rows) []string {
	columns := rows.FieldDescriptions()
	names := make([]string, len(columns))
	for i, column := range columns {
		names[i] = string(column.Name)
	}
	return names
}

// ScanStruct is
func ScanStruct(rows pgx.Rows, dest interface{}) error {
	m := map[string]interface{}{}
	err := ScanMap(rows, &m)
	if err != nil {
		return errors.Wrap(err, "Map")
	}
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   dest,
		TagName:  "db",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return errors.Wrap(err, "Decoder")
	}
	return decoder.Decode(m)
}
