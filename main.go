package main

import (
	"expense_log/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Создаем слайс пользователей для примера
var expTypes = []types.ExpenseType{
	{ID: uuid.New(), Name: "Groceries"},
	{ID: uuid.New(), Name: "Clothes"},
}

func main() {
	// Создаем новый роутер Gin
	r := gin.Default()

	// Эндпоинт для получения списка пользователей
	r.GET("/expType", func(c *gin.Context) {
		c.JSON(http.StatusOK, expTypes)
	})

	// Эндпоинт для добавления нового пользователя
	r.POST("/expType", func(c *gin.Context) {
		var newExpType types.ExpenseType

		// Парсим JSON из тела запроса в структуру newUser
		if err := c.ShouldBindJSON(&newExpType); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Присваиваем новый ID и добавляем пользователя в слайс
		newExpType.ID = uuid.New()
		// users = append(users, newUser)

		c.JSON(http.StatusCreated, newExpType)
	})

	// Запуск сервера на порту 8080
	r.Run(":8080")
}
