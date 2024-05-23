package db

import "github.com/jmoiron/sqlx"

type ExpenseStore struct {
	DB *sqlx.DB
}

func NewExpenseStore(dataSourceName string) (*ExpenseTypeStore, error) {
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &ExpenseTypeStore{DB: db}, nil
}
