package user

import (
	"e-commerce/domain/user"
	storageUser "e-commerce/infrastructure/postgres/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

func buildHandler(dbPool *pgxpool.Pool) handler {
	storage := storageUser.New(dbPool)
	useCase := user.New(storage)
	return newHandler(useCase)
}

