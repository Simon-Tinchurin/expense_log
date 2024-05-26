package main

import (
	"database/sql"
	"expense_log/types"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	dbname, dbuser, dbpass, dbhost, dbport, expTypeTableName, expensesTable string
)

func seedExpenseTypes(db *sql.DB) error {
	expenseTypes := []types.ExpenseType{
		{ID: uuid.New(), Name: "Groceries"},
		{ID: uuid.New(), Name: "Bars/Cafes"},
		{ID: uuid.New(), Name: "Liquor"},
		{ID: uuid.New(), Name: "Internet"},
		{ID: uuid.New(), Name: "Transport"},
		{ID: uuid.New(), Name: "Fastfood"},
		{ID: uuid.New(), Name: "Cigarettes"},
		{ID: uuid.New(), Name: "Delivery"},
		{ID: uuid.New(), Name: "Pharmacy"},
		{ID: uuid.New(), Name: "One-time costs"},
		{ID: uuid.New(), Name: "Financial costs"},
		{ID: uuid.New(), Name: "Clothes/Shoes"},
		{ID: uuid.New(), Name: "Rent"},
		{ID: uuid.New(), Name: "Utilities"},
		{ID: uuid.New(), Name: "Healthcare"},
	}

	for _, expenseType := range expenseTypes {
		// Check if the expense type already exists
		var exists bool
		err := db.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE name=$1)", expTypeTableName), expenseType.Name).Scan(&exists)
		if err != nil {
			return err
		}

		if !exists {
			_, err = db.Exec(fmt.Sprintf("INSERT INTO %s (id, name) VALUES ($1, $2)", expTypeTableName), expenseType.ID, expenseType.Name)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

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
	expTypeTableName = os.Getenv("EXPENSE_TYPES_TABLE")
	expensesTable = os.Getenv("EXPENSES_TABLE")
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

	// Creating ExpenseType table only if it doesn't exist
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id UUID PRIMARY KEY,
			name VARCHAR(255) NOT NULL
		);
	`, expTypeTableName)

	_, err = db.Exec(query)

	if err != nil {
		panic(err)
	}
	fmt.Println("ExpenseType table created successfully!")

	// Seeding the expense types
	err = seedExpenseTypes(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("Expense types seeded successfully!")

	// Creating Expenses table only if it doesn't exist
	query = fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id UUID PRIMARY KEY,
			date INT NOT NULL,
			expense_type_id UUID REFERENCES %s(id),
			price FLOAT NOT NULL,
			comment TEXT
		);
	`, expensesTable, expTypeTableName)

	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Expenses table created successfully!")
}
