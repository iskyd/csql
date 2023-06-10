package parser

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseInsertQuery(query string) (table string, columns []string, values []string, err error) {
	// Regular expression pattern to match the INSERT query structure
	pattern := `INSERT\s+INTO\s+(.*#*\w+)\s*\(([\w,\s-]+)\)\s+VALUES\s*\(([0-9\,\s\'A-Za-z]+)\);`

	// Compile the regular expression pattern
	re := regexp.MustCompile(pattern)

	// Find matches in the query using the regular expression
	matches := re.FindStringSubmatch(query)

	if len(matches) != 4 {
		err = fmt.Errorf("invalid INSERT query format")
		return
	}

	// Extract table name, column names, and values from the matches
	table = matches[1]
	columns = strings.Split(matches[2], ",")
	valuesRaw := strings.Split(matches[3], ",")

	// Trim whitespace from column names and values
	for i := range columns {
		columns[i] = strings.TrimSpace(columns[i])
	}
	for i := range valuesRaw {
		valuesRaw[i] = strings.TrimSpace(valuesRaw[i])
	}

	// Convert values to the appropriate types (you may need to handle different types based on your needs)
	for _, value := range valuesRaw {
		// If the value is enclosed in single quotes, treat it as a string
		if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
			values = append(values, strings.Trim(value, "'"))
		} else {
			// Otherwise, treat it as an integer (you can handle other types similarly)
			values = append(values, value)
		}
	}

	return
}
