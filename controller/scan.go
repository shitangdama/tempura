package controller

import (
	"database/sql"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// // FieldNames is
// func FieldNames(rows pgx.Rows) []string {
// 	columns := rows.FieldDescriptions()
// 	names := make([]string, len(columns))
// 	for i, column := range columns {
// 		names[i] = string(column.Name)
// 	}
// 	return names
// }

// ScanMap is
func ScanMap(rows *sql.Rows, dest *map[string]interface{}) error {
	// columns := FieldNames(rows)
	columns, err := rows.Columns()
	/*
		values := make([]interface{}, len(columns))
		for i := range values {
			x := new(interface{})
			values[i] = x
			//values[i] = "" //interface{}{}
		}
		err := rows.Scan(values...)
	*/
	if err != nil {
		return err
	}

	if !rows.Next() {
		return rows.Err()
	}

	values := make([]interface{}, len(columns))

	for index := range values {
		values[index] = new(sql.NullString)
	}

	err = rows.Scan(values...)

	if err != nil {
		return err
	}

	for index := range values {
		// values[index] = new(sql.NullString)
		(*dest)[string(column)] = values[i]
	}
	// values, err := rows.Values()
	// if err != nil {
	// 	return errors.Wrap(err, "ScanMap")
	// }

	// for i, column := range columns {
	// 	(*dest)[string(column)] = values[i]
	// 	//		(*dest)[string(column)] = *(values[i].(*interface{}))
	// }
	return rows.Err()
}

// ScanStruct is
func ScanStruct(rows *sql.Rows, dest interface{}) error {
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
