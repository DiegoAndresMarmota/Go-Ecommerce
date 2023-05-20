package purchaseorder

import (
	"context"
	"e-commerce/infrastructure/postgres"
	"e-commerce/model"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const tableDetails = "ticketpurchase_details"

var fieldsDetails = []string{
	"id",
	"invoice_id",
	"product_id",
	"amount",
	"unit_price",
	"created_at",
	"updated_at",
}

//BuildSQLInsert ingresa tableDetails y fieldsDetails, para que sea utilizado dentro de la interface de domain.TicketPurchase.Storage
var (
	psqlInsertDetails = postgres.BuildSQLInsert(tableDetails, fieldsDetails)
)

//CreateDetails recibe los detalles de la transacción
func (i TicketPurchase) CreateDetailsBulk(tx pgx.Tx, details model.TicketPurchaseDetails) error {
	batch := pgx.Batch{}
	for _, v := range details {
		batch.Queue(
			psqlInsertDetails,
			v.ID,
			v.TicketPurchaseID,
			v.ProductID,
			v.Amount,
			v.UnitPrice,
			v.CreatedAt,
			postgres.Int64ToNull(v.UpdatedAt),
		).Exec(func(ct pgconn.CommandTag) error {
			return nil
		})
	}

	result := tx.SendBatch(context.Background(), &batch)
	defer func() {
		err := result.Close()
		if err != nil {
			log.Printf("Error: %v", err)
		}
	}()

	_, err := result.Exec()
	if err != nil {
		return err
	}

	return nil
}