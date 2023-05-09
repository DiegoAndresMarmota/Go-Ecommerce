package user

import (
	"e-commerce/model"
	"time"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

//User representa la Storage
type User struct {
	storage Storage
}

//New crea un nuevo User, en el cual se almacena el storage de User
func New(s Storage) User {
	return User{storage: s}
}

//Creación de nuevo Usuario
func (u User) Create(m *model.User) error{
	ID, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("%s %w", "uuid.NewUUID()", err)
	}

	//Veriificación de ID
	m.ID = ID
	//Cifrado de password
	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s %w", "bcrypt.GenerateFromPassword()", err)
	}

	m.Password = string(password)
	if m.Details == nil {
		m.Details = []byte("{}")
	}
	m.CreatedAt = time.Now().Unix()

	err = u.storage.Create(m)
	if err != nil {
		return fmt.Errorf("%s %w", "storage.Create()", err)
	}

	m.Password = ""
	return nil
}

//GetByEmail recibe como parametro un email. Obteniendo a través de este parametro, la información contenida en el storage de User.
func (u User) GetByEmail(email string) (model.User, error) {
	user, err := u.storage.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("%s %w", "storage.GetByEmail()", err)
	}
	return user, nil
}

//GetAll recibe todos los registros almacenados en el storage de Users.
func (u User) GetAll() (model.Users, error) {
	users, err := u.storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("%s %w", "storage.GetAll()", err)
	}
	return users, nil
}