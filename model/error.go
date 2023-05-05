package model

import "fmt"

type Error struct {
	Code string
	Err error
	Who string
	StatusHTTP int
	Data any
	APIMessage string
	UserID string
}

func NewError() Error {
	return Error{}
}

//Declaración del metodo para el manejo de errores
func (e *Error) Error() string {
	return fmt.Sprintf("Code: %s, Err: %s, Who: %s, StatusHTTP: %d, Data: %v, UserID: %s",
		e.Code, e.Err, e.Who, e.StatusHTTP, e.Data, e.UserID,
	)
}

//Funcionalidades adicionales
func (e *Error) HasCode() bool {
	return e.Code != ""
}

func (e *Error) hasStatusHTTP() bool {
	return e.StatusHTTP > 0
}

func (e *Error) HasData() bool {
	return e.Data != nil
}
