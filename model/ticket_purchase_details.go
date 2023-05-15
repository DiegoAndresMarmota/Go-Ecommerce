package model

import (
	"github.com/google/uuid"
)

// TicketPurchaseDetail model of table TicketPurchaseDetails
type  TicketPurchaseDetail struct {
	ID uuid.UUID `json:"id"`
	TicketPurchaseID uuid.UUID `json:"invoice_id"`
	ProductID uuid.UUID `json:"product_id"`
	Amount    uint      `json:"amount"`
	UnitPrice float64   `json:"unit_price"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

//Verificación de TicketPurchase existe(ID)
func (i TicketPurchaseDetail) HasID() bool {
	return i.ID != uuid.Nil
}

// TicketPurchases slice of TicketPurchase
type TicketPurchaseDetails []TicketPurchaseDetail

//Verificación de TicketPurchase se encuentra vacío
func (i TicketPurchaseDetails) IsEmpty() bool { return len(i) == 0 }