package types

import "github.com/google/uuid"

type Expense struct {
	ID          uuid.UUID   `json:"id"`
	Date        int         `json:"date"`
	ExpenseType ExpenseType `json:"expenseType"`
	Price       float64     `json:"price"`
	Comment     string      `json:"comment"`
}

type ExpenseType struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
