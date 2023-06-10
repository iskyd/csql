package parser

import (
	"reflect"
	"testing"
)

func TestParseInsertQuery(t *testing.T) {
	query := "INSERT INTO .class (font-size, font-weight) VALUES (32, bold);"

	table, columns, values, err := ParseInsertQuery(query)

	// Assert that no error occurred
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Assert the expected values
	expectedTable := ".class"
	expectedColumns := []string{"font-size", "font-weight"}
	expectedValues := []string{"32", "bold"}

	if table != expectedTable {
		t.Errorf("Expected table: %s, got: %s", expectedTable, table)
	}

	if !reflect.DeepEqual(columns, expectedColumns) {
		t.Errorf("Expected columns: %v, got: %v", expectedColumns, columns)
	}

	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected values: %v, got: %v", expectedValues, values)
	}

	query = "INSERT INTO #id (font-size, font-weight) VALUES (32, bold);"
	expectedTable = "#id"

	table, columns, values, err = ParseInsertQuery(query)

	// Assert that no error occurred
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Assert the expected values
	if table != expectedTable {
		t.Errorf("Expected table: %s, got: %s", expectedTable, table)
	}

	if !reflect.DeepEqual(columns, expectedColumns) {
		t.Errorf("Expected columns: %v, got: %v", expectedColumns, columns)
	}
}
