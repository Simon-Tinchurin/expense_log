package types

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID          uuid.UUID   `json:"id"`
	Date        time.Time   `json:"date"`
	ExpenseType ExpenseType `json:"expense_type"`
	Price       float64     `json:"price"`
	Comment     string      `json:"comment"`
}

type ExpenseType struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
