package main

import (
	"expense_log/aggregator"
	"expense_log/db"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	dbname, dbuser, dbpass, dbport string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	dbname = os.Getenv("DB_NAME")
	dbuser = os.Getenv("POSTGRES_USER")
	dbpass = os.Getenv("POSTGRES_PASSWORD")
	dbport = os.Getenv("APP_PORT")
}

func main() {
	// Connecting to the DB
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbuser, dbpass, dbname)

	expTypeStore, err := db.NewExpenseTypesStore(dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	expStore, err := db.NewExpenseStore(dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()

	r.GET("/expType", aggregator.HandleGetExpType(expTypeStore))
	r.POST("/expType", aggregator.HandlePostExpType(expTypeStore))

	r.GET("/expTypes", aggregator.HandleGetExpTypes(expTypeStore))

	r.POST("/new-expense", aggregator.HandlePostExpense(expStore))
	r.GET("/get_expense/", aggregator.HandleGetExpenseByID(expStore))
	r.GET("/expenses-by-type", aggregator.HandleGetExpenseByType(expStore))

	r.Run(dbport)
}
