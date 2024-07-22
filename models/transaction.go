package models

import (
	"simpleBankingApplication/db"
	"strconv"
	"time"

	"golang.org/x/exp/rand"
)

type Transaction struct {
	ID                    int       `json:"id"`
	Amount                float64   `json:"amount" binding:"required"`
	TransactionType       string    `json:"transaction_type" binding:"required"`
	Transaction_reference string    `json:"transaction_reference"`
	AccountNumber         string    `json:"account_number" binding:"required"`
	UserID                int       `json:"user_id" binding:"required"`
	DateCreated           time.Time `json:"date_created"`
	DateModified          time.Time `json:"date_updated"`
}

func (transaction Transaction) SaveTransaction() error {
	transaction.Transaction_reference = GenerateRandomNumbers(30)
	_, err := db.DB.Exec("INSERT INTO transactions(amount,transaction_type,account_number,transaction_reference,user_id,date_created,date_updated) VALUES(?,?,?,?,?,?,?)", transaction.Amount, transaction.TransactionType, transaction.AccountNumber, transaction.Transaction_reference, transaction.UserID, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}

func GenerateRandomNumbers(n int) string {
	rand.Seed(uint64(time.Now().UnixNano())) // Initialize the random number generator.
	var randomNumberString string

	for i := 0; i < n; i++ {
		// Generate a random digit from 0 to 9 and append it to the string.
		randomNumberString += strconv.Itoa(rand.Intn(10))
	}

	return randomNumberString
}

func GetPaymentByReference(reference string) (Transaction, error) {
	var transaction Transaction
	row := db.DB.QueryRow("SELECT id,amount,transaction_type,transaction_reference,account_number,user_id,date_created,date_updated FROM transactions WHERE transaction_reference = ?", reference)
	err := row.Scan(&transaction.ID, &transaction.Amount, &transaction.TransactionType, &transaction.Transaction_reference, &transaction.AccountNumber, &transaction.UserID, &transaction.DateCreated, &transaction.DateModified)

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
