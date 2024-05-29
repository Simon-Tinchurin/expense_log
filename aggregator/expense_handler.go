package aggregator

import (
	"expense_log/db"
	"expense_log/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ExpenseRequest struct {
	Date            int     `json:"date"`
	ExpenseTypeName string  `json:"expense_type_name"`
	Price           float64 `json:"price"`
	Comment         string  `json:"comment"`
}

func HandlePostExpense(store *db.ExpenseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ExpenseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		expTypeId, err := store.GetExpenseTypeIdByName(req.ExpenseTypeName)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense type name"})
			return
		}

		newExpense := types.Expense{
			ID:   uuid.New(),
			Date: req.Date,
			ExpenseType: types.ExpenseType{
				ID:   expTypeId,
				Name: req.ExpenseTypeName,
			},
			Price:   req.Price,
			Comment: req.Comment,
		}

		if err := store.InsertExpense(newExpense); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, newExpense)
	}
}

// ExpenseIDRequest represents the request structure for retrieving an expense by ID
type ExpenseIDRequest struct {
	ID string `json:"id"`
}

// HandleGetExpenseByID handles POST requests to retrieve an expense by its ID from the request body
func HandleGetExpenseByID(store *db.ExpenseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ExpenseIDRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Parse the expense ID
		expenseID, err := uuid.Parse(req.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
			return
		}

		// Retrieve the expense from the store
		expense, err := store.GetExpense(expenseID.String())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve expense"})
			return
		}

		// Return the expense data as JSON
		c.JSON(http.StatusOK, expense)
	}
}
