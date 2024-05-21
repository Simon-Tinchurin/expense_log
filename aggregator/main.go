package aggregator

import (
	"database/sql"
	"expense_log/db"
	"expense_log/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type ExpenseTypeRequest struct {
	Name string `json:"name" binding:"required"`
}

func HandlePostExpType(store *db.ExpenseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ExpenseTypeRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check that Expense Type with this name does not exists in the DB
		if _, err := store.GetExpTypeByName(req.Name); err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Expense type with this name already exists"})
			return
		} else if err != sql.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		newExpenseType := types.ExpenseType{
			ID:   uuid.New(),
			Name: req.Name,
		}

		if err := store.InsertExpType(newExpenseType); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, newExpenseType)
	}
}

func HandleGetExpType(store *db.ExpenseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Name string `json:"name" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		expenseType, err := store.GetExpTypeByName(req.Name)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Expense type not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, expenseType)
	}
}
