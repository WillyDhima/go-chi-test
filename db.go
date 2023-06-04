package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB() (*sql.DB, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %w", err)
	}

	// Retrieve the database credentials from environment variables
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Verify that the environment variables are set
	if username == "" || password == "" || dbName == "" {
		return nil, fmt.Errorf("Missing database credential in environment variable")
	}

	// Construct the connection string
	dsn := fmt.Sprintf("%s:%s@%s", username, password, dbName)

	// Open Database Connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to the database: %w", err)
	}

	// Check if the connection is successfull
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error pinging to the database: %w", err)
	}

	log.Println("Connected to the database!")

	return db, nil
}
