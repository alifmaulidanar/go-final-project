package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Initializes the MySQL database (raw SQL)
func InitDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("MYSQL_PUBLIC_URL")
	if dsn == "" {
		log.Fatal("MYSQL_PUBLIC_URL environment variable is required but not set")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to the MySQL database")
	return db, nil
}
