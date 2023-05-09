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

//BuildSQLUpdate selectiona los campos identificados por ID y los actualiza
func BuildSQLUpdate(table string, fields []string) string {
	if len(fields) == 0 {
		return FieldEmpty
	}

	fields = append(fields[1:], fields[0])
	args := bytes.Buffer{}
	k := 0
	for _, v := range fields {
		if v == "created_at" {
			continue
		}
		args.WriteString(fmt.Sprintf("%s = $%d, ", v, k+1))
		k++
	}
	args.Truncate(args.Len() - 2)

	return fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", table, args.String(), k)
}

func BuildSQLDelete(table string) string {
	return fmt.Sprintf("DELETE FROM %s WHERE id = $1", table)
}