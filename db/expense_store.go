package db

import (
	"expense_log/types"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

type ExpenseStore struct {
	DB           *sqlx.DB
	ExpenseTable string
}

func NewExpenseStore(dataSourceName string) (*ExpenseStore, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file")
	}

	tableName := os.Getenv("EXPENSES_TABLE")
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &ExpenseStore{DB: db, ExpenseTable: tableName}, nil
}

func (store *ExpenseStore) InsertExpense(expense types.Expense) error {
	query := fmt.Sprintf("INSERT INTO %s (id, date, expense_type_id, price, comment) VALUES ($1, $2, $3, $4, $5)", pq.QuoteIdentifier(store.ExpenseTable))
	_, err := store.DB.Exec(query, expense.ID, expense.Date, expense.ExpenseTypeId, expense.Price, expense.Comment)
	return err
}

func (store *ExpenseStore) GetExpense(id string) (types.Expense, error) {
	var expense types.Expense
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", store.ExpenseTable)
	err := store.DB.QueryRow(query, id).Scan(&expense.ID, &expense.Date, &expense.ExpenseTypeId,
		&expense.Price, &expense.Comment)
	if err != nil {
		panic(err)
	}
	return expense, nil
}
