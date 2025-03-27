package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var db *sql.DB

func ConnectDB() (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
		return nil, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort, "disable")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("No se pudo verificar la conexión a la base de datos:", err)
		return nil, err
	}

	fmt.Println("Conexión exitosa a la base de datos")
	return db, nil
}
