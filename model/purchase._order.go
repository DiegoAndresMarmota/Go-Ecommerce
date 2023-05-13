package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

// PurchaseOrder estructura el tipo de TABLE de purchase_order
type PurchaseOrder struct {
	ID        uuid.UUID      `json:"id"`
	UserID    uuid.UUID       `json:"user_id"`
	Products  json.RawMessage `json:"products"`
	CreatedAt int64           `json:"created_at"`
	UpdatedAt int64           `json:"updated_at"`
}

//HasID es la validación de ID.PurchaseOrden
func (p PurchaseOrder) HasID() bool {
	return p.ID != uuid.Nil
}