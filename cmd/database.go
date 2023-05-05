package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/jackc/pgx/v5/pgxpool"
)

const AppCommerce = "Go-ECommerce"

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

//[] de conexión
func makeConnection (
	user, pass, host, port, dbName, sslMode string, minConn, maxConn int
	) string {
		return fmt.Sprintf("user=$s password=$s host=$s port=$s dbname=$s sslmode=$s pool_min_conns=$d pool_max_conns=%d",
		user,
		pass,
		host,
		port,
		dbName,
		sslMode,
		minConn,
		maxConn,
	)
}

	//connPrincipal, verifica y valida la conexión
	connPrincipal := makeConnection(user, pass, host, port, dbName, sslMode, min, max)
	config, err := pgxpool.ParseConfig(connPrincipal)
	if err != nil {
		nil, fmt.Errorf("error parsing config: %w", err, "%s", "pgxpool.ParseConfig()")
	}

	config.ConnConfig.RuntimeParams["aplication_name"] = AppCommerce

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		nil, fmt.Errorf("error parsing config: %w", err, "%s", "pgxpool.NewWithConfig()")
	}
	return pool, nil
}

