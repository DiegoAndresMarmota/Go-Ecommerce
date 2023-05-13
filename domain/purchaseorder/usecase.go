package purchaseorder

import (
	"e-commerce/model"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// PurchaseOrder implementara los métodos del Storage
type PurchaseOrder struct {
	storage Storage
}

// New realiza a nueva PurchaseOrder, según los métodos del Storage
func New(s Storage) PurchaseOrder {
	return PurchaseOrder{storage: s}
	}

// Create ingresa una nueva variable en model.PurchaseOrder
func (p PurchaseOrder) Create(m *model.PurchaseOrder) error {
	//Si m no es válido
	if err := m.Validate(); err != nil {
		return fmt.Errorf("purchaseorder: %w", err)
	}

	//Si hay error en ID
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	//Asignar time en el m de ID
	m.ID = ID
	m.CreatedAt = time.Now().Unix()

	//
	err = p.storage.Create(m)
	if err != nil {
		return err
	}

	return nil
}

//GetByID consulta y entrega el ID almacenado en el Storage, según la orden de pago
func (p PurchaseOrder) GetByID(ID uuid.UUID) (model.PurchaseOrder, error) {
	purchaseOrder, err := p.storage.GetByID(ID)
	if err != nil {
		return model.PurchaseOrder{}, fmt.Errorf("purchaseorder: %w", err)
	}

	return *purchaseOrder, nil
}