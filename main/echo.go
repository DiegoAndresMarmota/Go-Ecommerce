package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//Configuración de Echo
func newHTTP(errorHandler echo.HTTPErrorHandler) *echo.Echo {
	//Nueva instancia de echo
	e := echo.New()

	//Utilización de middleware [manejar peticiones y recuperar errores imprevistos]
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
}))

	e.HTTPErrorHandler = errorHandler

	return e
}
