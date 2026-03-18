package utils

import "fmt"

func MapFromArgs(args ...any) map[string]any {
	result := make(map[string]any)

	for i := 0; i < len(args); i += 2 {
		// 1. Safety check: make sure there is a value for the last key
		if i+1 >= len(args) {
			// Handle odd number of arguments (Key without a Value)
			if key, ok := args[i].(string); ok {
				result[key] = nil
			}
			break
		}

		// 2. Extract Key (must be a string)
		key, ok := args[i].(string)
		if !ok {
			// If the key isn't a string, we can force it to a string
			key = fmt.Sprintf("%v", args[i])
		}

		// 3. Extract Value
		result[key] = args[i+1]
	}

	return result
}
