package main

import (
	"os"
	"strings"
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

	corsConfig := middleware.CorsConfig{
		AllowOrigins: strings.Split(os.Getenv( key:"ALLOWED_ORIGINS"), sep: ","),
		ALLowMethods: strings.Split(os.Getenv( key:"ALLOWED_METHODS"), sep: ",")
	}

	e.Use(middleware.CORSWithConfig(corsConfig))

	e.HTTPErrorHandler = errorHandler

	return e
}
