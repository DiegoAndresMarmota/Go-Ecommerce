package purchaseorder

import "e-commerce/infrastructure/postgres"

const (
	table = "purchaseorder"
)

var fields = []string{
	"id",
	"user_id",
	"purchase_order_id",
	"created_at",
	"updated_at",
}

var (
	psqlInsert = postgres.BuildSQLInsert(table, fields)
)