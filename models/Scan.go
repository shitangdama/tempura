package models

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// // ScanMap is
// func ScanMap(rows pgx.Rows, dest *map[string]interface{}) error {
// 	columns := FieldNames(rows)
// 	/*
// 		values := make([]interface{}, len(columns))
// 		for i := range values {
// 			x := new(interface{})
// 			values[i] = x
// 			//values[i] = "" //interface{}{}
// 		}
// 		err := rows.Scan(values...)
// 	*/
// 	values, err := rows.Values()
// 	if err != nil {
// 		return errors.Wrap(err, "ScanMap")
// 	}

// 	for i, column := range columns {
// 		(*dest)[string(column)] = values[i]
// 		//		(*dest)[string(column)] = *(values[i].(*interface{}))
// 	}
// 	return rows.Err()
// }

// ScanMap is
func ScanMap(rows *sql.Rows, dest *map[string]interface{}) error {

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	values := make([]interface{}, len(columns))

	for index := range values {
		values[index] = new(sql.NullString)
	}

	if err := rows.Scan(values...); err != nil {
		return err
	}

	for index, columnName := range columns {
		(*dest)[string(columnName)] = *(values[index].(*sql.NullString))

		fmt.Println((*values[index].(*sql.NullString)).String)
		fmt.Println(reflect.ValueOf(values[index]).Elem())
	}
	return nil
}

// ScanMapArray is
func ScanMapArray() error {

	return nil
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
