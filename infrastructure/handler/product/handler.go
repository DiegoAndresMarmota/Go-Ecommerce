package product

import (
	"github.com/diegoandresmarmota/go-ecommerce/domain/product"

	"github.com/diegoandresmarmota/go-ecommerce/infrastructure/handler/response"

	"github.com/diegoandresmarmota/go-ecommerce/model"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase  product.UserCase
	response response.API
}

func newHandler(useCase product.UserCase) handler {
	return handler{useCase: useCase}
}

// Create Handler, maneja la creacion de model.Product
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

// Delete handler, maneja la eliminación del model.Product
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

// GetByID Handler, maneja le obtención de ID del model.Product
func (h handler) GetByID(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.Error(c, "uuid.Parse()", err)
	}

	productData, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.response.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.response.OK(productData))
}

// GetAll Handler, maneja la busqueda completa de todos los model.Product
func (h handler) GetAll(c echo.Context) error {
	products, err := h.useCase.GetAll()
	if err != nil {
		return h.response.Error(c, "useCase.GetAllWhere()", err)
	}

	return c.JSON(h.response.OK(products))
}
