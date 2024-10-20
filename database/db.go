package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func Initialize() error {
	var err error
	db, err = ConnectToDatabase()
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return fmt.Errorf("Failed to connect to the database: %w", err)
	}
	log.Printf("Connected to the MySQL Database.")

	return nil
}

func GetDB() *sql.DB {
	return db
}

func ConnectToDatabase() (*sql.DB, error) {
	host := os.Getenv("MYSQL_HOST")
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	port := os.Getenv("MYSQL_PORT")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/spread"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return nil, fmt.Errorf("ERROR: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Printf("ERROR: %v", err)
		return nil, fmt.Errorf("ERROR: %v", err)
	}

	return db, nil
}
