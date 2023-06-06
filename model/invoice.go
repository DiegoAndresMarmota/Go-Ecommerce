package model

import (
	"github.com/google/uuid"
)

// Invoice model of table Invoice
type Invoice struct {
	ID              uuid.UUID `json:"id"`
	UserID          uuid.UUID `json:"user_id"`
	PurchaseOrderID uuid.UUID `json:"purchase_order_id"`
	CreatedAt       int64     `json:"created_at"`
	UpdatedAt       int64     `json:"updated_at"`
}

//Verificación de Invoice existe(ID)
func (i Invoice) HasID() bool {
	return i.ID != uuid.Nil
}

// Invoices slice of TicketPurchase
type Invoices []Invoice

//Verificación de Invoice se encuentra vacío
func (i Invoices) IsEmpty() bool { return len(i) == 0 }

//Historial del slice de TicketPurchase, retornando el ID y el TicketPurchase asociado.
func (i Invoices) IDs() []uuid.UUID {
	var resp []uuid.UUID
	for _, v := range i {
		resp = append(resp, v.ID)
	}

	return resp
}