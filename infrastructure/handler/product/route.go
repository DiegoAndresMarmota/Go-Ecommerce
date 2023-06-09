package product

import (

	"github.com/diegoandresmarmota/go-ecommerce/infrastructure/handler/middleware"
	productStorage "github.com/diegoandresmarmota/go-ecommerce/infrastructure/postgres/product"

	"github.com/diegoandresmarmota/go-ecommerce/domain/product"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

// NewRouter devuelve una ruta para las peticiones efectuadas en el model.Product
func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	authMiddleware := middleware.New()

	//Se incorpora el middleware de Auth
	adminRoutes(e, h, authMiddleware.IsAdmin, authMiddleware.IsValid)
	publicRoutes(e, h)
}

func buildHandler(dbPool *pgxpool.Pool) handler {
	useCase := product.New(productStorage.New(dbPool))
	return newHandler(useCase)
}

// adminRoutes handle, maneja las rutas
func adminRoutes(e *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := e.Group("/api/v1/admin/products", middlewares...)

	route.POST("", h.Create)
	route.PUT("/:id", h.Update)
	route.DELETE("/:id", h.Delete)

	route.GET("", h.GetAll)
	route.GET("/:id", h.GetByID)
}

// publicRoutes maneja las rutas que no requieren de algún tipo de validación
func publicRoutes(e *echo.Echo, h handler) {
	route := e.Group("/api/v1/public/products")

	route.GET("", h.GetAll)
	route.GET("/:id", h.GetByID)
}
