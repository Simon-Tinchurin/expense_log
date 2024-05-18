package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	dbname, dbuser, dbpass, dbhost, dbport string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	dbname = os.Getenv("DB_NAME")
	dbuser = os.Getenv("POSTGRES_USER")
	dbpass = os.Getenv("POSTGRES_PASSWORD")
	dbhost = os.Getenv("DB_HOST")
	dbport = os.Getenv("DB_PORT")
}

func main() {
	// Creating a string for connecting to the PostgreSQL DB
	psqlInfo := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", dbname, dbuser, dbpass, dbhost, dbport)

	// Connecting to the PostgreSQL DB
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Checking connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	// Installing extension to automatically create id
	// Creating ExpenseType table
	_, err = db.Exec(`
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

		CREATE TABLE expense_types (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			name VARCHAR(255) NOT NULL
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("ExpenseType table created successfully!")
}
