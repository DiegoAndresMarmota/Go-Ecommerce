package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"database/sql"
	"e-commerce/infrastructure/handler"
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

	//Migrar SQL
	sqlDB, err := sql.Open("pgx")
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})


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

	handler.InitRoutes(e, dbPool)
	_ = dbPool

	//Inicialización
	err = e.Start(":" + os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}


//Pendiente Crear newDBConnection en database
//Pendiente Crear response en handleerror
