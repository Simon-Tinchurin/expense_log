package db

import (
	"database/sql"
	"expense_log/types"

	"github.com/jmoiron/sqlx"
)

type ExpenseTypeStore struct {
	DB *sqlx.DB
}

func NewExpenseStore(dataSourceName string) (*ExpenseTypeStore, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &ExpenseTypeStore{DB: db}, nil
}

func (store *ExpenseTypeStore) InsertExpType(expType types.ExpenseType) error {
	query := `INSERT INTO expense_types (id, name) VALUES ($1, $2)`
	_, err := store.DB.Exec(query, expType.ID, expType.Name)
	return err
}

func (store *ExpenseTypeStore) GetExpTypeByName(name string) (types.ExpenseType, error) {
	var expenseType types.ExpenseType
	query := "SELECT id, name FROM expense_types WHERE name = $1"
	err := store.DB.QueryRow(query, name).Scan(&expenseType.ID, &expenseType.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.ExpenseType{}, err
		}
		return types.ExpenseType{}, err
	}
	return expenseType, nil
}
