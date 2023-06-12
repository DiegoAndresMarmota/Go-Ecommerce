package purchaseorder

import (
	"context"
	"fmt"

	"github.com/diegoandresmarmota/go-ecommerce/infrastructure/postgres"

	"github.com/diegoandresmarmota/go-ecommerce/model"

	"github.com/jackc/pgx/v5"
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

// TicketPurchase struct implementa la interface de domain.TicketPurchase.Storage
type TicketPurchase struct {
	db *pgxpool.Pool
}

// New regresa a new TicketPurchase storage
func New(db *pgxpool.Pool) TicketPurchase {
	return TicketPurchase{db: db}
}

// $
func (i TicketPurchase) getTx() (pgx.Tx, error) {
	return i.db.Begin(context.Background())
}

// Create creates a model.TicketPurchase
func (i TicketPurchase) Create(m *model.TicketPurchase, ms model.TicketPurchaseDetails) error {
	tx, err := i.getTx()
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.UserID,
		m.PurchaseOrderID,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		errRollback := tx.Rollback(context.Background())
		if errRollback != nil {
			return fmt.Errorf("%s %w", errRollback, err)
		}

		return err
	}

	err = i.CreateDetailsBulk(tx, ms)
	if err != nil {
		errRollback := tx.Rollback(context.Background())
		if errRollback != nil {
			return fmt.Errorf("%s %w", errRollback, err)
		}

		return err
	}

	errCommit := tx.Commit(context.Background())
	if errCommit != nil {
		return errCommit
	}

	return nil
}
