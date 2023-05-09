package product

import (
	"e-commerce/model"
	"fmt"
	"github.com/google/uuid"
)

//Product implementa la estructura de Usercase de Storage
type Product struct {
	storage Storage
}

//New implementa la estructura de Storage en Product
func New(s Storage) Product {
	return Product{storage: s}
}

//Creación de un model.Product, a través de un ID(uuid)
func (p Product) Create(m *model.Product) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		fmt.Errorf("%v %d", "Error creating ID", "uuid.NewUUID()", err)
	}

	//Si el campo m.Images || m.Features de Product se encuentra vacío, almacenar un [] vacio
	m.ID = ID
	if len(m.Images) == 0 {
		m.Images = []byte(`[]`)
	}
	if len(m.Features) == 0 {
		m.Features = []byte(`[]`)
	}

	err = p.storage.Create(m)
	if err != nil {
		fmt.Errorf("%v %d", "Error creating field in model.Product", err)

	}

	return nil

}