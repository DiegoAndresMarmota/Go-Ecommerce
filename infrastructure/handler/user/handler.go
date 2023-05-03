package user

import (
	"e-commerce/domain/user"
	"e-commerce/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase user.UserCase
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
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, m)
}