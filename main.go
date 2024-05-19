package main

import (
	"expense_log/types"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var expTypes = []types.ExpenseType{
	{ID: uuid.New(), Name: "Groceries"},
	{ID: uuid.New(), Name: "Clothes"},
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
}

func main() {
	port := os.Getenv("APP_PORT")

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

	r.Run(port)
}
