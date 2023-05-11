package product

import (
	"e-commerce/model"
	"fmt"
	"time"

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

//Create crea p Product en el model.Product, a través de un ID(uuid)
func (p Product) Create(m *model.Product) error {
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%v %d", "Error creating ID", "uuid.NewUUID()", err)
	}

	//Si el campo m.Images || m.Features de Product se encuentra vacío, almacenar un [] vacio
	m.ID = ID
	if len(m.Images) == 0 {
		m.Images = []byte(`[]`)
	}
	if len(m.Features) == 0 {
		m.Features = []byte(`{}`)
	}

	err = p.storage.Create(m)
	if err != nil {
		return fmt.Errorf("%v %d", "Error creating field in model.Product", err)

	}

	return nil

}

//Update actualiza el model.Product, a través de la validación de .HasID
func(p Product) Update(m *model.Product) error {
	if !m.HasID() {
		return fmt.Errorf("ID product: %w", model.NewError().Err)
	}

	//Si el campo m.Images || m.Features de Product se encuentra vacío, almacenar un [] vacio
	if len(m.Images) == 0 {
		m.Images = []byte(`[]`)
	}
	if len(m.Features) == 0 {
		m.Features = []byte(`{}`)
	}
	//Añadir al m de model.Product, una hora local actual
	m.UpdatedAt = time.Now().Unix()

	//En caso de no actualizar el m de model.Product, enviar mensaje de error
	err := p.storage.Update(m)
	if err != nil {
		return err
	}

	return nil

}

// Delete elimina el model.Product por ID en el storage
func (p Product) Delete(ID uuid.UUID) error {
	err := p.storage.Delete(ID)
	if err != nil {
		return err
	}

	return nil
}

//GetByID retorna el model.Product segun el ID almacenado en el storage
func (p Product) GetByID(ID uuid.UUID) (model.Product, error) {
	product, err := p.storage.GetByID(ID)
	if err != nil {
		return model.Product{}, fmt.Errorf("product: %w", err)
	}

	return product, nil

}

//GetAll retorna products
func (p Product) GetAll() (model.Products, error) {
	products, err := p.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("product: %w", err)
	}

	return products, nil

}