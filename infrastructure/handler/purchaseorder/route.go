package purchaseorder

import (
	purchaseorderStorage "github.com/diegoandresmarmota/go-ecommerce/infrastructure/postgres/purchaseorder"

	"github.com/diegoandresmarmota/go-ecommerce/domain/purchaseorder"

	"github.com/diegoandresmarmota/go-ecommerce/infrastructure/handler/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
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
