package product

import (
	"e-commerce/model"

	"github.com/google/uuid"
)


type UserCase interface {
	Create(m *model.Product) error
	Update(m *model.Product) error
	Delete(ID uuid.UUID) error
	GetByID(ID uuid.UUID) (model.Product, error)
	GetAll() (model.Products, error)
}

type Storage interface {
	Create(m *model.Product) error
	Update(m *model.Product) error
	Delete(ID uuid.UUID) error
	GetByID(ID uuid.UUID) (model.Product, error)
	GetAll() (model.Products, error)
}