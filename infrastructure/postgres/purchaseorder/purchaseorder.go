package purchaseorder

import (
	"e-commerce/infrastructure/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

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

//TicketPurchase struct implementa la interface de domain.TicketPurchase.Storage
type TicketPurchase struct {
	db *pgxpool.Pool
}

//New regresa a new TicketPurchase storage
func New(db *pgxpool.Pool) TicketPurchase {
	return TicketPurchase{db: db}
}

