package models

import "simpleBankingApplication/db"

type User struct {
	ID       int    `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	AccountNumber string `json:"account_number" binding:"required"`
	AccountBalance float64 `json:"account_balance"`
}

func (user User) SaveUser() error {
	user.AccountBalance = 1000000.00 // this is the initial account balance
	_, err := db.DB.Exec("INSERT INTO users(first_name,last_name,account_number,account_balance) VALUES(?,?,?,?)", user.FirstName, user.LastName, user.AccountNumber,user.AccountBalance)

	if err != nil {
		return err
	}

	return nil
}

func CreditAccount(accountNumber string, amount float64) error { // this is for crediting the account
	_, err := db.DB.Exec("UPDATE users SET account_balance = account_balance + ? WHERE account_number = ?", amount, accountNumber)

	if err != nil {
		return err
	}

	return nil
}

func DebitAccount(id int, amount float64) error { // this is for debiting the account
	_, err := db.DB.Exec("UPDATE users SET account_balance = account_balance - ? WHERE id = ?", amount, id)

	if err != nil {
		return err
	}

	return nil
}
