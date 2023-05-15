package ticketpurchase

import (
	"e-commerce/model"
)

type UseCase interface {
	Create(m *model.PurchaseOrder) error
}

type Storage interface {
	Create(m *model.TicketPurchase, ms model.TicketPurchaseDetails) error
}