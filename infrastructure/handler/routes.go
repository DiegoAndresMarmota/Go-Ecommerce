package handler

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

//Inicializador de rutas
func InitRoutes(e *echo.Echo, dbPool *pgxpool.Pool) {
	health(e)
}

//Ruta Ok
func health(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
			return c.JSON(
				http.StatusOK,
				map[string]string{
					"time": time.Now().String(),
					"message": "OK",
					"service_name": "",
				},
			)
		})
	}