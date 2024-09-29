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
	// Load .env file only in local environment
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file, falling back to system environment variables")
		}
	}

	// Get the MySQL DSN from environment variables
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
