package user

import (
	"e-commerce/domain/user"
	"e-commerce/infrastructure/handler/response"
	"e-commerce/model"

	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase user.UserCase
	responsed response.API
}

// NewHandler retorna la implementación del dominio como un punto de entrada
func newHandler(uc user.UserCase) handler {
	return handler{useCase: uc}
}

//Crea la conexión con Echo
func (h *handler) Create(c echo.Context) (err error) {
	m := model.User{}

	//Si model.User == error, informar al cliente de la petición erronea con un .JSON del error
	if err := c.Bind(&m); err != nil {
		return h.responsed.BindFailed(err)
		// return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	//Devolución de Error estandar.
	if err := h.useCase.Create(&m); err != nil {
		return h.responsed.Error(c, "useCase.Create()", err)
		// return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(h.responsed.Created(m))
}

//Obtiene la información completa del userCase
func (h *handler) GetAll(c echo.Context) (err error) {
	users, err := h.useCase.GetAll()
	if err != nil {
		return h.responsed.Error(c, "useCase.GetAll()", err)
		// c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(h.responsed.OK(users))
}
