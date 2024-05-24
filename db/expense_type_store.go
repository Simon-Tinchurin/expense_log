package db

import (
	"database/sql"
	"expense_log/types"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type ExpenseTypeStore struct {
	DB            *sqlx.DB
	ExpTypesTable string
}

func NewExpenseTypesStore(dataSourceName string) (*ExpenseTypeStore, error) {
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

func (store *ExpenseTypeStore) InsertExpType(expType types.ExpenseType) error {
	query := fmt.Sprintf("INSERT INTO %s (id, name) VALUES ($1, $2)", store.ExpTypesTable)
	_, err := store.DB.Exec(query, expType.ID, expType.Name)
	return err
}

func (store *ExpenseTypeStore) GetExpTypeByName(name string) (types.ExpenseType, error) {
	var expenseType types.ExpenseType
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE name = $1", store.ExpTypesTable)
	err := store.DB.QueryRow(query, name).Scan(&expenseType.ID, &expenseType.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.ExpenseType{}, err
		}
		return types.ExpenseType{}, err
	}
	return expenseType, nil
}

func (store *ExpenseTypeStore) GetAllExpTypes() ([]types.ExpenseType, error) {
	var expTypes []types.ExpenseType
	query := fmt.Sprintf("SELECT * FROM %s", store.ExpTypesTable)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var expType types.ExpenseType
		if err := rows.Scan(&expType.ID, &expType.Name); err != nil {
			return nil, err
		}
		expTypes = append(expTypes, expType)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return expTypes, nil
}
