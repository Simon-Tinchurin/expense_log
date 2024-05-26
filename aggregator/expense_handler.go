package aggregator

import (
	"expense_log/db"
	"expense_log/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ExpenseRequest struct {
	Date          int     `json:"date"`
	ExpenseTypeId string  `json:"expense_type_id"`
	Price         float64 `json:"price"`
	Comment       string  `json:"comment"`
}

func HandlePostExpense(store *db.ExpenseStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ExpenseRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		expTypeId, err := uuid.Parse(req.ExpenseTypeId)
		if err != nil {
			panic(err)
		}

		newExpense := types.Expense{
			ID:            uuid.New(),
			Date:          req.Date,
			ExpenseTypeId: expTypeId,
			Price:         req.Price,
			Comment:       req.Comment,
		}
		fmt.Println(newExpense)
		if err := store.InsertExpense(newExpense); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, newExpense)
	}
}
