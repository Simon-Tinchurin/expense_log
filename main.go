package main

import (
	"expense_log/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var expTypes = []types.ExpenseType{
	{ID: uuid.New(), Name: "Groceries"},
	{ID: uuid.New(), Name: "Clothes"},
}

func main() {
	r := gin.Default()

	r.GET("/expType", func(c *gin.Context) {
		c.JSON(http.StatusOK, expTypes)
	})

	r.POST("/expType", func(c *gin.Context) {
		var newExpType types.ExpenseType

		if err := c.ShouldBindJSON(&newExpType); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newExpType.ID = uuid.New()

		c.JSON(http.StatusCreated, newExpType)
	})

	r.Run(":8080")
}
