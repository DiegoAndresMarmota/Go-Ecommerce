package ticketpurchase

import (
	"e-commerce/model"
	"fmt"
)

// TicketPurchase implements UseCase
type TicketPurchase struct {
	storage Storage
}

// New retorna un nuevo TicketPurchase
func New(s Storage) TicketPurchase {
	return TicketPurchase{storage: s}
}


// Create ingresa po a model.TicketPurchase
func (i TicketPurchase) Create(po *model.PurchaseOrder) error {
	//Validar si TicketPurchase es correcta
	if err := po.Validate(); err != nil {
		return fmt.Errorf("TicketPurchase: %w", err)
	}

	//Rellenar datos dentro de TicketPurchase && TicketPurchaseDetails en TicketPurchaseFromPurchaseOrder
	TicketPurchase, TicketPurchaseDetails, err := TicketPurchaseFromPurchaseOrder(po)
	if err != nil {
		return fmt.Errorf("%s %w", "TicketPurchaseFromPurchaseOrder()", err)
	}

	err = i.storage.Create(&TicketPurchase, TicketPurchaseDetails)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create()", err)
	}

	return nil
}


//Pendiente TicketPurchaseFromPurchaseOrder