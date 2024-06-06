package utils

import (
	"fmt"
	"strings"
)

func CompareSQLValues(sqlInsert, keyValues string) bool {
	// Parse the SQL insert command
	columnsStart := strings.Index(sqlInsert, "(")
	columnsEnd := strings.Index(sqlInsert, ")")
	columns := strings.Split(sqlInsert[columnsStart+1:columnsEnd], ",")

	valuesStart := strings.Index(sqlInsert, "VALUES")
	valuesStart += 6 // to skip "VALUES"
	valuesEnd := strings.LastIndex(sqlInsert, ")")
	values := strings.Split(sqlInsert[valuesStart+1:valuesEnd], ",")

	// Convert key-value pairs to map
	kvPairs := make(map[string]string)
	pairs := strings.Split(keyValues, ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		kvPairs[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
	}

	// Check if all provided columns have corresponding values
	for _, column := range columns {
		column = strings.TrimSpace(column)
		if value, ok := kvPairs[column]; !ok {
			fmt.Printf("Column '%s' not found in key values.\n", column)
			return false
		} else {
			// Check if the provided value matches the expected value
			if value != values[0] {
				fmt.Printf("Value '%s' for column '%s' does not match expected value.\n", value, column)
				return false
			}
			// Remove the value once checked
			values = values[1:]
		}
	}

	return true
}
