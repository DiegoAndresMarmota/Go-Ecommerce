package main

import (
	"log"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

//newDBConnection permite conectar a través de pgxpool con PostgreSQL
func newDBConnection() (*pgxpool.Pool, error) {
	min := 3
	max := 100

	minConn := os.Getenv(key: "DB_MIN_CONN")
	maxConn := os.Getenv(key: "DB_MAX_CONN")

	//Variables de entorno
	user := os.Getenv(key: "DB_USER")
	pass := os.Getenv(key: "DB_PASSWORD")
	host := os.Getenv(key: "DB_HOST")
	port := os.Getenv(key: "DB_PORT")
	dbName := os.Getenv(key: "DB_NAME")
	sslMode := os.Getenv(key: "DB_SSL_MODE")

	//Validación min && max
	if minConn != "" {
		v, err := strconv.Atoi(minConn)
		if err != nil {
			log.Printf("Error: DB_MIN_CONN has not a valid value", min)
		} else {
			if v >= 3 && v <= 100 {
				min = v
			}
		}
	}
	if maxConn != "" {
		v, err := strconv.Atoi(maxConn)
		if err != nil {
			log.Printf("Error: DB_MAX_CONN has not a valid value", max)
		} else {
			if v >= 3 && v <= 100 {
				max = v
			}
		}
	}
