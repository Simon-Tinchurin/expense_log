package db

import "expense_log/types"

type ExpenseStore interface {
	GetExpById()
	InsertExp(expTypeName types.ExpenseType)
	UpdateExp(ID int)
	DeleteExp(ID int)
}
