package postgres

import (
	"database/sql"
	"bytes"
	"fmt"
)

var FieldEmpty = "The field is empty"

//BuildSQLInsert crea el nombre de la tabla y los campos(argumento y valor) que se van a rellenar
func BuildSQLInsert(table string, fields []string) string {
	if len(fields) == 0 {
		return FieldEmpty

	args := bytes.Buffer{}
	values := bytes.Buffer{}
	k:= 0

	for _, v := range fields {
		k++
		args.WriteString(v)
		args.WriteString( ", ")
		values.WriteString(fmt.Sprintf("$#{k},"))
	}

	args.Truncate(args.Len() - 2)
	values.Truncate(values.Len() - 2)

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, args.String(), values.String())

	}
