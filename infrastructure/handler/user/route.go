package user

import (
	"e-commerce/domain/user"
	"e-commerce/infrastructure/handler/middleware"
	storageUser "e-commerce/infrastructure/postgres/user"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

//Ruta Principal
func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middleware.New()
	adminRoutes(e, h, authMiddleware.IsAdmin, authMiddleware.IsValid)
	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	storage := storageUser.New(dbPool)
	useCase := user.New(storage)
	return newHandler(useCase)
}

//Ruta de admin
func adminRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	g := e.Group("/api/v1/admin/users", middlewares...)

	g.GET("", h.GetAll)
}

func publicRoutes(e *echo.Echo, h handler) {
	g := e.Group("/api/v1/public/users")

	g.POST("", h.Create)
}
