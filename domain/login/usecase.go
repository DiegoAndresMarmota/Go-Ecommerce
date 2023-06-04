package login

import (
	"e-commerce/model"
	"fmt"
	"time"
	"github.com/golang-jwt/jwt"
)

// Login implementa su struct
type Login struct {
	useCaseUser UseCaseUser
}

//New retorna un Login de UseCaseUser de su struct
func New(uc UseCaseUser) Login {
	return Login{useCaseUser: uc}
}

//Login recibe el email, password y jwtSecretKey, valida y lo entrega al user. Se crea el model.JWT, donde el token expira en 2 horas, donde este tiene un metodo de firma HS256, enviandole el structSecreKey, y este firme el token(data). Devolviendo el user y la data firmada.
func (l Login) Login(email, password, jwtSecretKey string) (model.User, string, error) {
	user, err := l.useCaseUser.Login(email, password)
	if err != nil {
		return model.User{}, "", fmt.Errorf("%s %w", "useCaseUser.Login()", err)
	}

	structSecreKey := model.JWT{
		UserID:  user.ID,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, structSecreKey)

	data, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return model.User{}, "", fmt.Errorf("%s %w", "token.SignedString()", err)
	}

	user.Password = ""

	return user, data, nil
}