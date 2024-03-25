package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DBConn represents the database connection
var DBConn *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	// Connection string for PostgreSQL
	connStr := "postgres://postgres:bhavya@22@localhost/GeoData?sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	// Set the database connection to the global variable
	DBConn = db

	fmt.Println("Connected to the database")

	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if DBConn != nil {
		DBConn.Close()
		fmt.Println("Database connection closed")
	}
}
