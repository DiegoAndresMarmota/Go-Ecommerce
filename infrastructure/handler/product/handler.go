package product

import (
	"e-commerce/domain/product"
	"e-commerce/infrastructure/handler/response"
	"e-commerce/model"

	"github.com/google/uuid"
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

// Update handler, maneja la actualización del model.Product
func (h handler) Update(c echo.Context) error {
	m := model.Product{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}
	m.ID = ID

	if err := h.useCase.Update(&m); err != nil {
		return h.response.Error(c, "useCase.Update()", err)
	}

	return c.JSON(h.response.Updated(m))
}

// Delete handles the deleting of a model.Product
func (h handler) Delete(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}

	err = h.useCase.Delete(ID)
	if err != nil {
		return h.response.Error(c, "useCase.Delete()", err)
	}

	return c.JSON(h.response.Deleted(nil))
}