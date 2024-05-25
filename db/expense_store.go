package db

import (
	"expense_log/types"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type ExpenseStore struct {
	DB           *sqlx.DB
	ExpenseTable string
}

func NewExpenseStore(dataSourceName string) (*ExpenseTypeStore, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file")
	}

	tableName := os.Getenv("EXPENSE_TYPES_TABLE")
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &ExpenseTypeStore{DB: db, ExpTypesTable: tableName}, nil
}

func (store *ExpenseStore) InsertExpense(expense types.Expense) error {
	query := fmt.Sprintf("INSERT INTO %s (id, date, expense_type_id, price, comment) VALUES ($1, $2)", store.ExpenseTable)
	_, err := store.DB.Exec(query, expense.ID, expense.Date, expense.ExpenseType.ID, expense.Price, expense.Comment)
	return err
}
