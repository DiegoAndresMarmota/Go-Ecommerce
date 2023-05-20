package main

import (
	"e-commerce/infrastructure/handler/response"
	"log"
	"os"
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
	err = e.Start(":" + os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}


//Pendiente Crear newDBConnection en database
//Pendiente Crear response en handleerror