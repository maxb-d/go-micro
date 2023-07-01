package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	ToAccount int `json:"to_account"`
	Amount    int `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
type Account struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"firstname"`
	LastName      string    `json:"lastname"`
	AccountNumber int64     `json:"account_number"`
	Balance       float64   `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
}

func NewAccount(firstname string, lastname string) *Account {
	return &Account{
		FirstName:     firstname,
		LastName:      lastname,
		AccountNumber: int64(rand.Intn(100000)),
		CreatedAt:     time.Now().UTC(),
	}
}
