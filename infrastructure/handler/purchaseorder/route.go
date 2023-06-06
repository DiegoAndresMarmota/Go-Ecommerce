package purchaseorder

import (
	"e-commerce/domain/purchaseorder"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	purchaseorderStorage "e-commerce/infrastructure/postgres/purchaseorder"
	"e-commerce/infrastructure/handler/middleware"
)

// NewRouter retorna un router para conectarse  a la petición de handle model.PurchaseOrder
func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middleware.New()
	privateRoutes(e, h, authMiddleware.IsValid)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCase := purchaseorder.New(purchaseorderStorage.New(dbPool))
	return newHandler(useCase)
}

// privateRoutes handler realiza la conexión mediante el token
func privateRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("/api/v1/private/purchase-orders", middlewares...)

	route.POST("", h.Create)
}