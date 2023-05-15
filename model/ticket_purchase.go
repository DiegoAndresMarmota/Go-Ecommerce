package model

import (
	"github.com/google/uuid"
)

// TicketPurchase model of table TicketPurchase
type TicketPurchase struct {
	ID              uuid.UUID `json:"id"`
	UserID          uuid.UUID `json:"user_id"`
	PurchaseOrderID uuid.UUID `json:"purchase_order_id"`
	CreatedAt       int64     `json:"created_at"`
	UpdatedAt       int64     `json:"updated_at"`
}

//Verificación de TicketPurchase existe(ID)
func (i TicketPurchase) HasID() bool {
	return i.ID != uuid.Nil
}

// TicketPurchases slice of TicketPurchase
type TicketPurchases []TicketPurchase

//Verificación de TicketPurchase se encuentra vacío
func (i TicketPurchases) IsEmpty() bool { return len(i) == 0 }

//Historial del slice de TicketPurchase, retornando el ID y el TicketPurchase asociado.
func (i TicketPurchases) IDs() []uuid.UUID {
	var resp []uuid.UUID
	for _, v := range i {
		resp = append(resp, v.ID)
	}

	return resp
}