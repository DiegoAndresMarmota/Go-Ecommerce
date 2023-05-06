package user

import (
	"context"
	"database/sql"
	"e-commerce/model"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

//Declaración de Crear Usuario y Obtener lista de usuario.
var (
	psqlInsert = "INSERT INTO users (id, email, password, details, created_at) VALUES ($1, $2, $3, $4, $5)"
	psqlGetAll = "SELECT id, email, password, details, created_at, updated_at FROM users"
)

type User struct {
	db *pgxpool.Pool
}

//Creación del CRUD con pxlpool en User
func New(db *pgxpool.Pool) *User {
	return &User{db: db}
}

//Creación de User a través de psqlInsert en el model.User
func (u *User) Create(m *model.User) error {
	_, err := u.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.Email,
		m.Password,
		m.IsAdmin,
		m.Details,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

//Obtención de Users a través de psqlSelect en los []model.User, añadiendo Where email
func (u User) GetByEmail(email string) (model.User, error) {
	query := psqlGetAll + "WHERE email = $1"
	row := u.db.QueryRow(
		context.Background(),
		query,
		email,
	)
	return u.scanRow(row, true)
}

//Retorna un []Users
func (u User) GetAll() (model.Users, error) {
	rows, err := u.db.Query(
		context.Background(),
		psqlGetAll,
	)

	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()



	ms:= model.Users{}
	for rows.Next() {
		m, err := u.scanRow(rows, false)
		if err != nil{
			ms = append(ms, m)
		}
	}
	return ms, nil
}


//Utiliza scanRow para obtener todos los registros de la base de datos para verificarlos y luego convertirlos en una estructura en el model.User
func (u User) scanRow(s pgx.Row, withPassword bool) (model.User, error) {
	m := model.User{}

	updatedAtNull := sql.NullInt64{}

	err:= s.Scan(
		&m.ID,
		&m.Email,
		&m.IsAdmin,
		&m.Details,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return m, err
	}

	m.UpdatedAt = updatedAtNull.Int64

	//Limpiar password
	if !withPassword {
		m.Password = ""
	}

	return m, nil
}


