package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Initializes the MySQL database (raw SQL)
func InitDB() (*sql.DB, error) {
	// Get the MySQL DSN from environment variables
	dsn := os.Getenv("MYSQL_PUBLIC_URL") // Ensure this is correctly set in Railway
	if dsn == "" {
		log.Fatal("MYSQL_PUBLIC_URL environment variable is required but not set")
	}

	// Add "tcp()" if not already present
	if !containsTcpWrapper(dsn) {
		dsn = addTcpWrapper(dsn)
	}

	// Connect to the MySQL database
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

// Helper function to ensure the DSN is using tcp()
func containsTcpWrapper(dsn string) bool {
	return len(dsn) >= 4 && dsn[0:4] == "tcp("
}

func addTcpWrapper(dsn string) string {
	return "@tcp(" + dsn + ")"
}
