package product

import (
	"e-commerce/infrastructure/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

//Table, como nombre de la tabla.
const table = "products"

//Fields, como los []de campos que tendra la tabla.
var fields = []string{
	"id",
	"product_name",
	"price",
	"images",
	"description",
	"features",
	"created_at",
	"updated_at",
}

//SQLBuilder para la tabla
var (
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlUpdate = postgres.BuildSQLUpdate(table, fields)
	psqlDelete = postgres.BuildSQLDelete(table, fields)
	psqlGetAll = postgres.BuildSQLGetAll(table, fields)
)

//Product struct, implementa la interface de domain.product.Storage, recibiendo pgxpool como conexión.
type Product struct {
	db *pgxpool.Pool
}

//New retorna un nuevo Product en el storage
func New(db *pgxpool.Pool) Product {
	return Product{db:db}
}

