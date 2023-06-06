package middleware

import (
	"e-commerce/infrastructure/handler/response"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)


type AuthMiddleware struct {
	responser response.API
}

func New() AuthMiddleware {
	return AuthMiddleware{}
}

// AuthMiddleware proporciona la validez del token con isValid
func (au AuthMiddleware) IsValid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := getTokenFromRequest(c.Request())
		if err != nil {
			return au.responser.BindFailed(err)
		}

		isValid, claims := au.validate(token)
		if !isValid {
			err = errors.New("Token is not valid")
			return au.responser.BindFailed(err)
		}

		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("isAdmin", claims.IsAdmin)

		return next(c)
	}
}

//getTokenFromRequest, recibe un request, obteniendolo del Header, entregando data, para chequear el token.
func getTokenFromRequest(r *http.Request) (string, error) {
	data := r.Header.Get("Authorization")
	if data == "" {
		return "", errors.New("Header is not valid")
	}

	if strings.HasPrefix(data, "Bearer") {
		return data[7:], nil
	}

	return data, nil
}