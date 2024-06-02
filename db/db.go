package db

type Store struct {
	ExpenseType ExpenseTypeStore
	Expense     ExpenseStore
}

type Pagination struct {
	Limit int64
	Page  int64
}
