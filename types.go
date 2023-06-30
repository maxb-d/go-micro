package main

import "math/rand"

type Account struct {
	ID            int     `json:"id"`
	FirstName     string  `json:"firstname"`
	LastName      string  `json:"lastname"`
	AccountNumber int64   `json:"account_number"`
	Balance       float64 `balance:"id"`
}

func NewAccount(firstname string, lastname string) *Account {
	return &Account{
		ID:            rand.Intn(10000),
		FirstName:     firstname,
		LastName:      lastname,
		AccountNumber: int64(rand.Intn(100000)),
	}
}
