package purchaseorder

import (

)

// PurchaseOrder implementara los métodos del Storage
type PurchaseOrder struct {
	storage Storage
}

// New realiza a nueva PurchaseOrder, según los métodos del Storage
func New(s Storage) PurchaseOrder {
	return PurchaseOrder{storage: s}
}