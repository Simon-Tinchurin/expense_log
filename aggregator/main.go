package main

import (
	"database/sql"
	"encoding/json"
	"expense_log/types"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	dbname = "your_db_name"
	dbuser = "your_db_user"
	dbpass = "your_db_password"
	dbhost = "your_db_host"
	dbport = "your_db_port"
)

func main() {
	http.HandleFunc("/create-expense-type", createExpenseTypeHandler)
	http.ListenAndServe(":8080", nil)
}

func createExpenseTypeHandler(w http.ResponseWriter, r *http.Request) {

	var expenseType types.ExpenseType
	err := json.NewDecoder(r.Body).Decode(&expenseType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	psqlInfo := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", dbname, dbuser, dbpass, dbhost, dbport)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO expense_types (name) VALUES ($1)", expenseType.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expenseType)
}
