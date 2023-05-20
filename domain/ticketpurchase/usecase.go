package ticketpurchase

import (
	"e-commerce/model"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
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
	ticketPurchase, ticketPurchaseDetails, err := ticketPurchaseFromPurchaseOrder(po)
	if err != nil {
		return fmt.Errorf("%s %w", "TicketPurchaseFromPurchaseOrder()", err)
	}

	err = i.storage.Create(&ticketPurchase, ticketPurchaseDetails)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create()", err)
	}

	return nil
}

//TicketPurchaseFromPurchaseOrder recibe una orden de compra y devuelve un ticket de la compra y el detalle del mismo, además de un error.
func ticketPurchaseFromPurchaseOrder(po *model.PurchaseOrder) (model.TicketPurchase, model.TicketPurchaseDetails, error) {
	ID, err := uuid.NewUUID()
	if err!= nil {
			return model.TicketPurchase{}, nil, fmt.Errorf("%s %w", "uuid.NewUUID()", err)
		}

	//Creación del ticket de compra
	ticketPurchase := model.TicketPurchase{
		ID: ID,
		UserID: po.UserID,
		PurchaseOrderID: po.ID,
		CreatedAt: time.Now().Unix(),
	}

	//Creación del detalle del ticket de compra
	var products model.ProductToPurchases
	err = json.Unmarshal(po.Products, &products)
	if err != nil {
		return model.TicketPurchase{}, nil, fmt.Errorf("%s %w", "json.Unmarshal()", err)
	}

	var ticketPurchaseDetails model.TicketPurchaseDetails
	for _, v := range products {
		detailID, err := uuid.NewUUID()
		if err!= nil {
					return model.TicketPurchase{}, nil, fmt.Errorf("%s %w", "json.Unmarshal()", err)
				}

		detail := model.TicketPurchaseDetail{
			ID: detailID,
			TicketPurchaseID: ticketPurchase.ID,
			ProductID: v.ProductID,
			Amount: v.Amount,
			UnitPrice: v.UnitPrice,
			CreatedAt: time.Now().Unix(),
		}

		ticketPurchaseDetails = append(ticketPurchaseDetails, detail)
	}

	return ticketPurchase, ticketPurchaseDetails, nil
}