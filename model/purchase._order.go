package model

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// PurchaseOrder es la estructura de TABLE
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

type PurchaseOrders []PurchaseOrder

type ProductToPurchases []ProductToPurchase

//ProductToPurchase es la estructura de TABLE, los sigtes campos
type ProductToPurchase struct {
	ProductID uuid.UUID `json:"product_id"`
	Amount    uint      `json:"amount"`
	UnitPrice float64   `json:"unit_price"`
}

//Validate chequea la validez(!=0) del json de Products
func (p PurchaseOrder) Validate() error {
	if len(p.Products) == 0 {
		return fmt.Errorf(" EL json de Products esta vacío")
	}

	//ProductToPurchase agregará p.Products en el shopping_cart
	var shopping []ProductToPurchase
	err := json.Unmarshal(p.Products, &shopping)
	if err != nil {
		return fmt.Errorf("%s %w", "json.Unmarshal()", err)
	}

	//Validación de cada campo de ProductToPurchase
	for _, v := range shopping {
		if v.ProductID == uuid.Nil {
			return fmt.Errorf("%s %w", "La ID del producto no se encuentra", err)
		}
		if v.Amount < 1 {
			return fmt.Errorf("%s %w", "La cantidad de productos debe ser mayor a uno", err)
		}
		if v.UnitPrice < 0 {
			return fmt.Errorf("%s %w", "El valor unitario del producto es erroneo", err)
		}
	}

	return nil
}

//TotalAmount establece si la cantidad de productos y el valor  total de cada uno de ellos.
func (p PurchaseOrder) TotalAmount() float64 {
	if len(p.Products) == 0 {
		return 0
	}

	//Validación de agregar p.Products en Shopping
	var shopping []ProductToPurchase
	err := json.Unmarshal(p.Products, &shopping)
	if err != nil {
		return 0
	}

	//Calculo del total de la orden de compra
	var total float64
	for _, v := range shopping {
		total += float64(v.Amount) * v.UnitPrice
	}

	return total
}

//IsEmpty retorna True si hay p en PurchaseOrders
func (p PurchaseOrders) IsEmpty() bool {
	return len(p) == 0
}