package db

import (
	"expense_log/types"
	"fmt"
	"os"

	"github.com/google/uuid"
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
	query := fmt.Sprintf("INSERT INTO %s (id, date, expense_type, price, comment) VALUES ($1, $2, $3, $4, $5)", pq.QuoteIdentifier(store.ExpenseTable))
	_, err := store.DB.Exec(query, expense.ID, expense.Date, expense.ExpenseType, expense.Price, expense.Comment)
	return err
}

// func (store *ExpenseStore) GetExpense(id string) (types.Expense, error) {
// 	var expense types.Expense
// 	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", store.ExpenseTable)
// 	err := store.DB.QueryRow(query, id).Scan(&expense.ID, &expense.Date, &expense.ExpenseTypeId,
// 		&expense.Price, &expense.Comment)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return expense, nil
// }

func (store *ExpenseStore) GetExpense(id string) (types.Expense, error) {
	var expense types.Expense
	query := fmt.Sprintf(`
        SELECT e.id, e.date, e.expense_type_id, e.price, e.comment,
               et.id, et.name
        FROM %s e
        JOIN expense_types et ON e.expense_type_id = et.id
        WHERE e.id = $1
    `, store.ExpenseTable)

	var expenseType types.ExpenseType
	err := store.DB.QueryRow(query, id).Scan(
		&expense.ID, &expense.Date, &expenseType.ID,
		&expense.Price, &expense.Comment,
		&expenseType.ID, &expenseType.Name,
	)
	if err != nil {
		return types.Expense{}, err
	}
	expense.ExpenseType = expenseType
	return expense, nil
}

func (store *ExpenseStore) GetExpenseTypeIdByName(name string) (uuid.UUID, error) {
	var id uuid.UUID
	query := "SELECT id FROM expense_types WHERE name = $1"
	err := store.DB.QueryRow(query, name).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id, nil
}
