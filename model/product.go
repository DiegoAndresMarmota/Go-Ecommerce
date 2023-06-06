package model

import (
	"encoding/json"
	"github.com/google/uuid"
)

//Product es la estructura del modelo para []products
type Product struct {
	ID          uuid.UUID       `json:"id"`
	ProductName string          `json:"product_name"`
	Price       float64         `json:"price"`
	Images      json.RawMessage `json:"images"`
	Description string          `json:"description"`
	Features    json.RawMessage `json:"features"`
	CreatedAt   int64           `json:"created_at"`
	UpdatedAt   int64           `json:"updated_at"`
}

//HasID es la validación de ID.product
func (p Product) HasID() bool {
	return p.ID != uuid.Nil
}

type Products []Product

//Is Empty es un validación si != 0 del []Product
func (p Products) IsEmpty() bool { return len(p) == 0 }