package purchaseorder

import (
	"github.com/diegoandresmarmota/go-ecommerce/model"

	"github.com/google/uuid"
)

// UseCase es la interfase que utilizara Create o GetByID para acceder en model.PurchaseOrder
type UseCase interface {
	Create(m *model.PurchaseOrder) error
	GetByID(ID uuid.UUID) (*model.PurchaseOrder, error)
}

// Storage es la interfase que utilizara Create o GetByID para acceder en model.PurchaseOrder
type Storage interface {
	Create(m *model.PurchaseOrder) error
	GetByID(ID uuid.UUID) (*model.PurchaseOrder, error)
}
