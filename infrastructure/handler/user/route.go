package user

import (

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	storageUser "e-commerce/infrastructure/postgres/user"
	"e-commerce/domain/user"
)

//Ruta Principal
func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	adminRoutes(e, h)
	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	storage := storageUser.New(dbPool)
	useCase := user.New(storage)

	return newHandler(useCase)
}

//Ruta de admin
func adminRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/admin/users")

	g.GET("", h.GetAll)
}

func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/public/users")

	g.POST("", h.Create)
}
