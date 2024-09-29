package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	// Get environment variables directly
	dbUser := os.Getenv("MYSQLUSER")
	dbPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbHost := os.Getenv("MYSQLHOST")     // Make sure this is set to the correct host
	dbPort := os.Getenv("MYSQLPORT")     // Ensure the correct port is provided
	dbName := os.Getenv("MYSQLDATABASE") // Confirm the database name

	// Check if any variables are empty
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		log.Fatal("One or more required environment variables are missing")
	}

	// Build the DSN string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to MySQL
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
