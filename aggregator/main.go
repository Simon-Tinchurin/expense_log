package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("POSTGRES_USER")
	DBPass := os.Getenv("POSTGRES_PASSWORD")
	fmt.Println("expense aggregator started")
	fmt.Printf("DB Name: %s; DB User: %s; DB Pass: %s\n", DBName, DBUser, DBPass)
}
