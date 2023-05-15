package model

import (
	"github.com/google/uuid"
)

// ticketPurchase model of table ticketPurchase
type ticketPurchase struct {
	ID              uuid.UUID `json:"id"`
	UserID          uuid.UUID `json:"user_id"`
	PurchaseOrderID uuid.UUID `json:"purchase_order_id"`
	CreatedAt       int64     `json:"created_at"`
	UpdatedAt       int64     `json:"updated_at"`
}

//Verificación de ticketPurchase existe(ID)
func (i ticketPurchase) HasID() bool {
	return i.ID != uuid.Nil
}

// ticketPurchases slice of ticketPurchase
type ticketPurchases []ticketPurchase

//Verificación de ticketPurchase se encuentra vacío
func (i ticketPurchases) IsEmpty() bool { return len(i) == 0 }

//Historial del slice de ticketPurchase, retornando el ID y el ticketPurchase asociado.
func (i ticketPurchases) IDs() []uuid.UUID {
	var resp []uuid.UUID
	for _, v := range i {
		resp = append(resp, v.ID)
	}

	return resp
}