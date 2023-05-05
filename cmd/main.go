package cmd

import (
	"log"
	"os"

	"github.com/DiegoAndresMarmota/go-ecommerce/infrastructure/handler/response"
)

func main() {
	//Verificar errores en el .file de variables de entorno
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	//Validar variables de entorno cargadas
	err = validateEnv()
	if err != nil {
		log.Fatal(err)
	}

	//Crear Echo --> Routing
	e := newHTTP(response.HTTPErrorHandler)

	//Instanciar conexión a base de datos
	dbPool, err := newDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	_ = dbPool

	//Inicialización
	err = e.Start(":" + os.Getenv(key: "SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
