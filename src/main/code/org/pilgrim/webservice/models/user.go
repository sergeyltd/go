package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/rs/xid"
)

type Transaction struct {
	ID        string
	Timestamp string
	Title     string
	Amount    float64
	Note      string
}

var (
	transactions []*Transaction
)

func GetNextId() string {
	return xid.New().String()
}

func GetTransactions() []*Transaction {
	return transactions
}

func AddTransaction(u Transaction) (Transaction, error) {
	if u.ID != "" {
		return Transaction{}, errors.New("New Transaction must not include id or it must be set to zero")
	}
	u.ID = GetNextId()
	u.Timestamp = time.Now().String()
	transactions = append(transactions, &u)
	return u, nil
}

func GetTransactionByID(id string) (Transaction, error) {
	for _, u := range transactions {
		if u.ID == id {
			return *u, nil
		}
	}

	return Transaction{}, fmt.Errorf("Transaction with ID '%v' not found", id)
}

func UpdateTransaction(u Transaction) (Transaction, error) {
	for i, candidate := range transactions {
		if candidate.ID == u.ID {
			transactions[i] = &u
			return u, nil
		}
	}

	return Transaction{}, fmt.Errorf("Transaction with ID '%v' not found", u.ID)
}

func RemoveTransactionById(id string) error {
	for i, u := range transactions {
		if u.ID == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Transaction with ID '%v' not found", id)
}
