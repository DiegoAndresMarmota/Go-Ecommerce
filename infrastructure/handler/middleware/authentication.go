package middleware

import (
	"e-commerce/infrastructure/handler/response"
	"e-commerce/model"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)


type AuthMiddleware struct {
	responsed response.API
}

func New() AuthMiddleware {
	return AuthMiddleware{}
}

// AuthMiddleware proporciona la validez del token con isValid
func (au AuthMiddleware) IsValid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := getTokenFromRequest(c.Request())
		if err != nil {
			return au.responsed.BindFailed(err)
		}

		isValid, claims := au.validate(token)
		if !isValid {
			err = errors.New("Token is not valid")
			return au.responsed.BindFailed(err)
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

//AuthMiddleware recibe de validate un token y entrega si es valido y el JWT
func (au AuthMiddleware) validate(token string) (bool, model.JWT) {
	claims, err := jwt.ParseWithClaims(token, &model.JWT{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		log.Println(token)
		log.Println(os.Getenv("JWT_SECRET_KEY"))
		log.Println(err)
		return false, model.JWT{}
	}

	data, ok := claims.Claims.(*model.JWT)
	if !ok {
		log.Println("is not a jwt valid")
		return false, model.JWT{}
	}

	return true, *data
}

//AuthMiddleware recibe IsAdmin, luego de validar si el user IsValid, recibiendo un echo.Context
func (am AuthMiddleware) IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin, ok := c.Get("isAdmin").(bool)
		if !isAdmin || !ok {
			err := errors.New("you are not admin")
			return am.responsed.BindFailed(err)
		}

		return next(c)
	}
}