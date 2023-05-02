package user

import (
	"e-commerce/model"
	"time"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	storage Storage
}

func New(s Storage) User {
	return User{storage: s}
}

//Creación de nuevo Usuario
func (u User) Create(m *model.User) {
	ID, err := uuid.NewUUID()
	if err != nil {
		return
		//fmt.Printf("%s %w", "uuid.NewUUID()", err) 
	}

	m.ID = ID
	//Cifrado de password
	password, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return
		//fmt.Printf("%s %w", "bcrypt.GenerateFromPassword()", err)
	}

	m.Password = string(password)
	if m.Details == nil {
		m.Details = []byte("{}")
	}
	m.CreatedAt = time.Now().Unix()

	err = u.storage.Create(m)
	if err != nil {
		return
		//fmt.Printf("%s %w", "storage Create()", err) 
	}

	m.Password = ""
	//return nil
}

//GetByEmail recibe como parametro un email. Obteniendo a traves de este parametro, la información contenida en el storage de User.
func (u User) GetByEmail(email string) (model.User, error) {
	user, err := u.storage.GetByEmail(email)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

//GetAll recibe todos los registros almacenados en el storage de Users.
func (u User) GetAll() (model.Users, error) {
	users, err := u.storage.GetAll()
	if err != nil {
		return model.Users{}, err
	}
	return users, nil
}