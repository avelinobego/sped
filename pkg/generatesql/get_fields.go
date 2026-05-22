package generatesql

import (
	"fmt"
	"reflect"
	"strings"
)

func getStructFields(model interface{}) ([]string, error) {
	t := reflect.TypeOf(model)

	// Dereference pointer if needed
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("model must be a struct, got %s", t.Kind())
	}

	var columns []string

	for field := range t.Fields() {

		tag := field.Tag.Get("db")
		if tag == "" || tag == "-" || tag == "id" {
			continue
		}

		colName := strings.Split(tag, ",")[0]
		columns = append(columns, colName)
	}

	if len(columns) == 0 {
		return nil, fmt.Errorf("no 'db' tags found on struct %s", t.Name())
	}

	return columns, nil
}
