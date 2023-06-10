package csql

import (
	"github.com/iskyd/csql/internal/converter"
	"github.com/iskyd/csql/internal/parser"
)

func Convert(sql string) (string, error) {
	table, columns, values, err := parser.ParseInsertQuery(sql)

	if err != nil {
		return "", err
	}

	return converter.Sql2Css(table, columns, values)
}
