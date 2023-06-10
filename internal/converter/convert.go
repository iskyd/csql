package converter

import "fmt"

func Sql2Css(table string, columns []string, values []string) (string, error) {
	css := table + " {\n"

	for i, column := range columns {
		css += fmt.Sprintf("\t%s: %s;\n", column, values[i])
	}

	css += "}"

	return css, nil
}
