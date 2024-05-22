package db

type Store struct {
	ExpenseType ExpenseTypeStore
}

type Pagination struct {
	Limit int64
	Page  int64
}
