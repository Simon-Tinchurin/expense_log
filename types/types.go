package types

import "github.com/google/uuid"

type Expense struct {
	ID          uuid.UUID   `json:"id"`
	Date        int         `json:"date"`
	ExpenseType ExpenseType `json:"expenseType"`
}

type ExpenseType struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
