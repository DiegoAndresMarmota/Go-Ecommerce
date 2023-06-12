package purchaseorder

import (

	"github.com/diegoandresmarmota/go-ecommerce/domain/purchaseorder"
	"github.com/diegoandresmarmota/go-ecommerce/infrastructure/handler/response""
	"fmt"

	"github.com/diegoandresmarmota/go-ecommerce/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase  purchaseorder.UseCase
	response response.API
}

func newHandler(useCase purchaseorder.UseCase) handler {
	return handler{useCase: useCase}
}

// Create-handler realiza la creación del m en el model.PurchaseOrder
func (h handler) Create(c echo.Context) error {
	m := model.PurchaseOrder{}
	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	userID, ok := c.Get("userID").(uuid.UUID)
	if !ok {
		return h.response.Error(c, "c.Get().(uuid.UUID)", fmt.Errorf("error"))
	}

	m.UserID = userID
	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.response.Created(m))
}
