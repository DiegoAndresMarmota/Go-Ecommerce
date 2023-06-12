package ticketpurchase

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/diegoandresmarmota/go-ecommerce/model"
	"github.com/google/uuid"
)

// Invoice implements UseCase
type Invoice struct {
	storage Storage
	storageInvoiceDetailReport StorageInvoiceDetailReport
}

// New retorna un nuevo Invoice
func New(s Storage, sidr StorageInvoiceDetailReport) Invoice {
	return Invoice{storage: s, storageInvoiceDetailReport: sidr}
}


// Create ingresa po a model.TicketPurchase
func (i Invoice) Create(po *model.PurchaseOrder) error {
	if err := po.Validate(); err != nil {
		return fmt.Errorf("invoice: %w", err)
	}

	invoice, invoiceDetails, err := invoiceFromPurchaseOrder(po)
	if err != nil {
		return fmt.Errorf("%s %w", "invoiceFromPurchaseOrder()", err)
	}

	err = i.storage.Create(&invoice, invoiceDetails)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create()", err)
	}

	return nil
}

//TicketPurchaseFromPurchaseOrder recibe una orden de compra y devuelve un ticket de la compra y el detalle del mismo, además de un error.
func invoiceFromPurchaseOrder(po *model.PurchaseOrder) (model.Invoice, model.InvoiceDetails, error) {
	ID, err := uuid.NewUUID()
	if err != nil {
		return model.Invoice{}, nil, fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	//Creación del ticket de compra
	invoice := model.Invoice{
		ID:              ID,
		UserID:          po.UserID,
		PurchaseOrderID: po.ID,
		CreatedAt:       time.Now().Unix(),
	}

	//Creación del detalle del ticket de compra
	var products model.ProductToPurchases
	err = json.Unmarshal(po.Products, &products)
	if err != nil {
		return model.Invoice{}, nil, fmt.Errorf("%s %w", "json.Unmarshal()", err)
	}

	var invoiceDetails model.InvoiceDetails
	for _, v := range products {
		detailID, err := uuid.NewUUID()
		if err != nil {
			return model.Invoice{}, nil, fmt.Errorf("%s %w", "uuid.NewUUID()", err)
		}

		detail := model.InvoiceDetail{
			ID:        detailID,
			InvoiceID: invoice.ID,
			ProductID: v.ProductID,
			Amount:    v.Amount,
			UnitPrice: v.UnitPrice,
			CreatedAt: time.Now().Unix(),
		}

		invoiceDetails = append(invoiceDetails, detail)
	}

	return invoice, invoiceDetails, nil
}