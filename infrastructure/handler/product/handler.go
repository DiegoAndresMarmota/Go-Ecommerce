package product

import (
	"e-commerce/domain/product"
	"e-commerce/infrastructure/handler/response"
	"e-commerce/model"

	"github.com/labstack/echo/v4"
)

//
type handler struct {
	useCase product.UserCase
	response response.API
}

//
func newHandler(useCase product.UserCase) handler {
	return handler{useCase: useCase}
}

//Create Handler, maneja la creacion de model.Product
func (h handler) Create(c echo.Context) error {
	m := model.Product{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	if err := h.useCase.Create(&m); err != nil {
		h.response.Error(c, "useCase.Create", err)
	}

	return c.JSON(h.response.Created(m))
}