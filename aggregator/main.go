package aggregator

import (
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
