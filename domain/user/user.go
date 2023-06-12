package user

import "github.com/diegoandresmarmota/go-ecommerce/model"

// Interface del dominio (ingreso de los datos de cliente)
type UserCase interface {
	Create(m *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() ([]model.User, error)
}

// Interface para almacenamiento (salida de los datos de cliente)
type Storage interface {
	Create(m *model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() ([]model.User, error)
	//GetAll() (model.Users, error) ***Optional
}
