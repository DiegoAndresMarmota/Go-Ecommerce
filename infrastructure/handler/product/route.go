package product

import (
	"e-commerce/domain/product"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	productStorage "e-commerce/infrastructure/postgres/product"
)

// NewRouter devuelve una ruta para las peticiones efectuadas en el model.Product
func NewRouter(e *echo.Echo, dbPool *pgxpool.Pool) {
	h := buildHandler(dbPool)

	adminRoutes(e, h)
	publicRoutes(e, h)
}

//
func buildHandler(dbPool *pgxpool.Pool) handler {
	userCase := product.New(productStorage.New(dbPool))
	return newHandler(userCase)
}


// adminRoutes handle, maneja las rutas
func adminRoutes(e *echo.Echo, h handler) {
	route := e.Group("/api/v1/admin/products")

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