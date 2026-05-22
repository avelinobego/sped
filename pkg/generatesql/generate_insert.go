package generatesql

import (
	"fmt"
	"strings"
)

// BuildInsert generates an INSERT SQL statement and returns the query string.
// It uses reflection to read `db` tags from the struct fields.
func BuildInsert(table string, model any) (string, error) {

	var params []string
	columns, err := getStructFields(model)
	if err != nil {
		return "", err
	}

	for _, col := range columns {
		params = append(params, ":"+col)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(columns, ", "),
		strings.Join(params, ", "),
	)

	return query, nil
}
