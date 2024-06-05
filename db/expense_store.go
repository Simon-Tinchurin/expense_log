package db

import (
	"expense_log/types"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

type ExpenseStore interface {
	InsertExpense(types.Expense) error
	GetExpense(string) (types.Expense, error)
	GetExpenseTypeIdByName(string) (uuid.UUID, error)
	GetExpensesByType(string) ([]types.Expense, error)
}

type PostgresExpStore struct {
	DB           *sqlx.DB
	ExpenseTable string
}

func NewExpenseStore(dataSourceName string) (*PostgresExpStore, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file")
	}

	tableName := os.Getenv("EXPENSES_TABLE")
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &PostgresExpStore{DB: db, ExpenseTable: tableName}, nil
}

func (store *PostgresExpStore) InsertExpense(expense types.Expense) error {
	query := fmt.Sprintf("INSERT INTO %s (id, date, expense_type_id, price, comment) VALUES ($1, $2, $3, $4, $5)", pq.QuoteIdentifier(store.ExpenseTable))
	_, err := store.DB.Exec(query, expense.ID, expense.Date, expense.ExpenseType.ID, expense.Price, expense.Comment)
	return err
}

func (store *PostgresExpStore) GetExpense(id string) (types.Expense, error) {
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

func (store *PostgresExpStore) GetExpenseTypeIdByName(name string) (uuid.UUID, error) {
	var id uuid.UUID
	query := "SELECT id FROM expense_types WHERE name = $1"
	err := store.DB.QueryRow(query, name).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id, nil
}

// GetExpensesByType retrieves expenses by their type name from the data source
func (store *PostgresExpStore) GetExpensesByType(typeName string) ([]types.Expense, error) {
	query := `
		SELECT e.id, e.date, e.expense_type_id, e.price, e.comment, et.name
		FROM expenses e
		JOIN expense_types et ON e.expense_type_id = et.id
		WHERE et.name = $1
	`
	rows, err := store.DB.Query(query, typeName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []types.Expense
	for rows.Next() {
		var expense types.Expense
		var expenseTypeName string
		if err := rows.Scan(&expense.ID, &expense.Date, &expense.ExpenseType.ID, &expense.Price, &expense.Comment, &expenseTypeName); err != nil {
			return nil, err
		}
		expense.ExpenseType.Name = expenseTypeName
		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

func (store *PostgresExpStore) GetMonthExpenses(month *int) ([]types.Expense, error) {
	var year, selectedMonth int
	currentYear, currentMonth, _ := time.Now().Date()

	if month != nil && *month >= 1 && *month <= 12 {
		selectedMonth = *month
		if selectedMonth > int(currentMonth) {
			year = currentYear - 1 // если выбранный месяц больше текущего, предполагаем, что это прошлый год
		} else {
			year = currentYear
		}
	} else {
		year = currentYear
		selectedMonth = int(currentMonth)
	}

	monthStart := time.Date(year, time.Month(selectedMonth), 1, 0, 0, 0, 0, time.UTC)
	monthEnd := monthStart.AddDate(0, 1, 0) // первый день следующего месяца

	query := `SELECT e.id, e.date, e.expense_type_id, e.price, e.comment, et.id, et.name
	FROM expenses e
	JOIN expense_types et ON e.expense_type_id = et.id
	WHERE e.date >= $1 AND e.date < $2`

	rows, err := store.DB.Query(query, monthStart, monthEnd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []types.Expense
	for rows.Next() {
		var expense types.Expense
		var expenseTypeID uuid.UUID
		var expenseTypeName string
		err := rows.Scan(
			&expense.ID,
			&expense.Date,
			&expenseTypeID,
			&expense.Price,
			&expense.Comment,
			&expenseTypeID,
			&expenseTypeName,
		)
		if err != nil {
			return nil, err
		}
		expense.ExpenseType = types.ExpenseType{ID: expenseTypeID, Name: expenseTypeName}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}
