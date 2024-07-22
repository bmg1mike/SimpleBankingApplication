package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb(){
	var err error
	DB, err = sql.Open("sqlite3","banking.db")

	if err != nil {
		panic("error connecting to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTables()
}

func CreateTables(){
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT,
		last_name TEXT,
		account_number TEXT,
		account_balance float64
	)
	`

	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("error creating user table")
	}

	createTransactionTable := `
	Create TABLE IF NOT EXISTS transactions(
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		amount REAL,
		transaction_type TEXT,
		transaction_reference TEXT,
		user_id INTEGER,
		account_number TEXT,
		date_created DATETIME,
		date_updated DATETIME
	)
	`

	_, err = DB.Exec(createTransactionTable)

	if err != nil {
		panic("error creating transaction table")
	}
}