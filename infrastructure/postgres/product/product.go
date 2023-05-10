package product

import (
	"context"
	"e-commerce/infrastructure/postgres"
	"e-commerce/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"database/sql"

)

//Table, como nombre de la tabla.
const table = "products"

//Fields, como los []de campos que tendra la tabla.
var fields = []string{
	"id",
	"product_name",
	"price",
	"images",
	"description",
	"features",
	"created_at",
	"updated_at",
}

//SQLBuilder para la tabla
var (
	psqlInsert = postgres.BuildSQLInsert(table, fields)
	psqlUpdate = postgres.BuildSQLUpdate(table, fields)
	psqlDelete = postgres.BuildSQLDelete(table)
	psqlGetAll = postgres.BuildSQLGetAll(table, fields)
)

//Product struct, implementa la interface de domain.product.Storage, recibiendo pgxpool como conexión.
type Product struct {
	db *pgxpool.Pool
}

//New retorna un nuevo Product en el storage
func New(db *pgxpool.Pool) Product {
	return Product{db:db}
}

//Create, crea a model.Product, .exec ingresa y rellena los datos de los datos correspondientes con el m de model.Product
func (p Product) Create(m *model.Product) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlInsert,
		m.ID,
		m.ProductName,
		m.Price,
		m.Images,
		m.Description,
		m.Features,
		m.CreatedAt,
		postgres.Int64ToNull(m.UpdatedAt),
	)
	if err != nil {
		return err
	}

	return nil
}

//Update, actualiza el model.Product, .exec edita los datos correspondientes con el m de model.Product
func (p Product) Update(m *model.Product) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlUpdate,
		m.ProductName,
		m.Price,
		m.Images,
		m.Description,
		m.Features,
		m.UpdatedAt,
		m.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

//Delete, elimina el model.Product según ID del Product
func (p Product) Delete(ID uuid.UUID) error {
	_, err := p.db.Exec(
		context.Background(),
		psqlDelete,
		ID,
	)
	if err != nil {
		return err
	}

	return nil

}

//GetByID obtiene un model.Product, según ID.
func (p Product) GetByID(ID uuid.UUID) (model.Product, error) {
	query := psqlGetAll + " WHERE id = $1"
	row := p.db.QueryRow(
		context.Background(),
		query,
		ID,
	)

	return p.scanRow(row)
}

//GetAll obtiene todos los model.Products con sus campos respectivos
func (p Product) GetAll() (model.Products, error) {
	rows, err := p.db.Query(
		context.Background(),
		psqlGetAll,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ms model.Products
	for rows.Next() {
		m, err := p.scanRow(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}

	return ms, nil

}

//scanRow retorna un scan de pgx.Row de model.Product
func (p Product) scanRow(s pgx.Row) (model.Product, error) {
	m := model.Product{}

	updatedAtNull := sql.NullInt64{}

	err := s.Scan(
		&m.ID,
		&m.ProductName,
		&m.Price,
		&m.Images,
		&m.Description,
		&m.Features,
		&m.CreatedAt,
		&updatedAtNull,
	)
	if err != nil {
		return m, err
	}

	m.UpdatedAt = updatedAtNull.Int64

	return m, nil
}