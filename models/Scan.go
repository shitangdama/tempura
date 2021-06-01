package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx/reflectx"
	"github.com/pkg/errors"
)

// ScanMap is
func ScanMap(rows *sql.Rows, dest *map[string]interface{}) error {

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	values := make([]interface{}, len(columns))

	for index := range values {
		// values[index] = new(sql.NullString)
		values[index] = new(interface{})
	}

	if err := rows.Scan(values...); err != nil {
		return err
	}

	for index, columnName := range columns {
		(*dest)[string(columnName)] = *values[index].(*interface{})

		fmt.Println(*values[index].(*interface{}))

	}
	return nil
}

var mapper = reflectx.NewMapperFunc("db", strings.ToLower)

// https://github.com/jmoiron/sqlx/blob/a1d5e64734233358bc4e0a54beddc18509071a95/sqlx.go#L895
// ScanStruct is
func ScanStruct(rows *sql.Rows, dest interface{}) error {

	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value, to ScanStruct destination")
	}
	if v.IsNil() {
		return errors.New("nil pointer passed to ScanStruct destination")
	}

	base := v.Type()
	if base.Kind() == reflect.Ptr {
		base = base.Elem()
	}

	columns, err := rows.Columns()

	if err != nil {
		return err
	}
	fields := mapper.TraversalsByName(v.Type(), columns)
	fmt.Println(44444)
	fmt.Println(columns)
	fmt.Println(v.Type())
	fmt.Println(fields)
	fmt.Println(44444)
	// if we are not unsafe and are missing fields, return an error
	if f, err := missingFields(fields); err != nil {
		fmt.Println(3333)
		return fmt.Errorf("missing destination name %s in %T", columns[f], dest)
	}
	values := make([]interface{}, len(columns))

	err = fieldsByTraversal(v, fields, values)
	fmt.Println(22222)
	if err != nil {
		return err
	}
	fmt.Println(11111)
	fmt.Println(v)
	fmt.Println(fields)
	fmt.Println(values)
	// scan into the struct field pointers and append to our results
	return rows.Scan(values...)
}

func missingFields(transversals [][]int) (field int, err error) {
	for i, t := range transversals {
		if len(t) == 0 {
			return i, errors.New("missing field")
		}
	}
	return 0, nil
}

func fieldsByTraversal(v reflect.Value, traversals [][]int, values []interface{}) error {
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return errors.New("argument is not a struct")
	}

	for i, traversal := range traversals {
		if len(traversal) == 0 {
			values[i] = new(interface{})
			continue
		}

		f := reflectx.FieldByIndexes(v, traversal)
		if f.Kind() == reflect.Ptr {
			values[i] = f.Interface()
		} else {
			values[i] = f.Addr().Interface()
		}
	}

	return nil
}
