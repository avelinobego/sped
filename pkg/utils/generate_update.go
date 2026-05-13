package utils

import (
	"fmt"
	"strings"
)

// BuildUpdate generates an SQL UPDATE statement for the given table and struct.
// The struct is used to determine the columns to update, and the idField is used to specify the WHERE clause.
func BuildUpdate(tableName string, model interface{}, idField string) (string, error) {
	columns, err := getStructFields(model)
	if err != nil {
		return "", err
	}

	setClauses := make([]string, 0)
	for _, col := range columns {
		// if col != idField {
		setClauses = append(setClauses, col+" = :"+col)
		// }
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = :%s", tableName, strings.Join(setClauses, ", "), idField, idField)
	return query, nil
}
