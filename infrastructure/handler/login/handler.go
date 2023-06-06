package login

import (
	"database/sql"
	"e-commerce/domain/login"
	"e-commerce/infrastructure/handler/response"
	"e-commerce/model"
	"errors"
	"net/http"
	"os"
	"strings"
	"github.com/labstack/echo/v4"
)

type handler struct {
	useCase   login.UseCase
	responsed response.API
}

//newHandler instancia un nuevo handler
func newHandler(useCase login.UseCase) handler {
	return handler{useCase: useCase}
}

//
func (h handler) Login(c echo.Context) error {
	m := model.Login{}
	err := c.Bind(&m)
	if err != nil {
		return h.responsed.BindFailed(err)
	}

	u, t, err := h.useCase.Login(m.Email, m.Password, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		if strings.Contains(err.Error(), "bcrypt.CompareHashAndPassword()") ||
			errors.Is(err, sql.ErrNoRows) {
			resp := model.MessageResponse{
				Data:     "wrong user or password",
				Messages: model.Responses{{Code: response.AuthError, Message: "wrong user or password"}},
			}
			return c.JSON(http.StatusBadRequest, resp)
		}
		return h.responsed.Error(c, "useCase.Login()", err)
	}

	return c.JSON(h.responsed.OK(map[string]interface{}{"user": u, "token": t}))
}