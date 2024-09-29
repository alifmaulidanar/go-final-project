package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Initializes the MySQL database (raw SQL)
func InitDB() (*sql.DB, error) {
	// Get the MySQL DSN from environment variables (in URL format from Railway)
	url := os.Getenv("MYSQL_PUBLIC_URL")
	if url == "" {
		log.Fatal("MYSQL_PUBLIC_URL environment variable is required but not set")
	}

	// Convert to Go's MySQL DSN format
	dsn := convertToMySQLDSN(url)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to the MySQL database")
	return db, nil
}

// Converts the Railway connection string to Go's MySQL DSN format
func convertToMySQLDSN(url string) string {
	url = url[len("mysql://"):]
	dsn := url[:len(url)-1]
	dsn = dsn[:len(dsn)-1] + "@tcp(" + dsn[len(dsn)-1:]

	return dsn
}
