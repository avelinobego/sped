package utils

import (
	"fmt"
	"reflect"
	"strings"
)

// BuildInsert generates an INSERT SQL statement and returns the query string.
// It uses reflection to read `db` tags from the struct fields.
func BuildInsert(table string, model any) (string, error) {
	t := reflect.TypeOf(model)

	// Dereference pointer if needed
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("model must be a struct, got %s", t.Kind())
	}

	var columns []string
	var params []string

	for i := range t.NumField() {
		field := t.Field(i)

		tag := field.Tag.Get("db")
		if tag == "" || tag == "-" {
			continue
		}

		// Support "column,omitempty" style tags — take only the name part
		colName := strings.Split(tag, ",")[0]

		columns = append(columns, colName)
		params = append(params, ":"+colName) // sqlx named parameter style
	}

	if len(columns) == 0 {
		return "", fmt.Errorf("no 'db' tags found on struct %s", t.Name())
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(columns, ", "),
		strings.Join(params, ", "),
	)

	return query, nil
}
