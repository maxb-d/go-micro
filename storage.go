package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Interfaces is a collection of function signatures
type Storage interface {
	GetCollectionAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccountByID(int) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=bandofdummiesdb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS Accounts(
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		account_number serial,
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	return err
}

func (s *PostgresStore) GetCollectionAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM ACCOUNTS")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	rows, err := s.db.Query("SELECT * FROM ACCOUNTS WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `INSERT INTO 
		Accounts (first_name, last_name, account_number, balance, created_at) 
		VALUES ($1, $2, $3, $4, $5)
	`
	resp, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.AccountNumber,
		acc.Balance,
		acc.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccountByID(id int) error {
	_, err := s.db.Query("DELETE FROM ACCOUNTS WHERE id = $1", id)

	return err
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.AccountNumber,
		&account.Balance,
		&account.CreatedAt)

	return account, err
}
