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

	// remove "`", "'" like extra characters
	for i, col := range columns {
		columns[i] = removeExtraChar(&col)
	}

	valuesStart := strings.Index(sqlInsert, "VALUES")
	valuesStart += 7 // to skip "VALUES"
	valuesEnd := strings.LastIndex(sqlInsert, ")")
	values := strings.Split(sqlInsert[valuesStart+1:valuesEnd], ",")

	// Convert key-value pairs to map
	kvPairs := make(map[string]string)
	pairs := strings.Split(keyValues, ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, ":")
		kvPairs[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
	}

	// Check if all provided columns have corresponding values
	for i, column := range columns {
		column = strings.TrimSpace(column)
		if value, ok := kvPairs[column]; !ok {
			fmt.Printf("Column '%s' not found in key values.\n", column)
		} else {
			// Check if the provided value matches the expected value
			parvalue := strings.TrimSpace(removeExtraChar(&values[i]))
			if value != parvalue {
				fmt.Printf("Column '%s' does not match. Expected value '%s' but received '%s'.\n", column, parvalue, value)
				return false
			}
		}
	}

	return true
}

func removeExtraChar(col *string) string {
	parsedCol := ""
	for _, currStr := range *col {
		switch string(currStr) {
		case "`", "'", `"`:
			break
		default:
			parsedCol += string(currStr)
		}
	}
	return parsedCol
}
