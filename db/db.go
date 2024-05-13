package db

type Store struct {
	Expense ExpenseStore
}

type Pagination struct {
	Limit int64
	Page  int64
}
