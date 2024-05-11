package types

type Expense struct {
	ID          int         `json:"id"`
	Date        int         `json:"date"`
	ExpenseType ExpenseType `json:"expenseType"`
}

type ExpenseType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
