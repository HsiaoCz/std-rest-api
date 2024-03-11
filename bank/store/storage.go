package store

import (
	"database/sql"

	"github.com/HsiaoCz/std-rest-api/bank/types"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*types.Account) error
	DeleteAccount(int) error
	UpdateAccount(*types.Account) error
	GetAccountByID(int) (*types.Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=verify-full"
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
	str := `CREATE TABLE IF NOT EXISTS account (
		 id serial primary key,
		 first_name varchar(50),
		 last_name varchar(50),
		 number serial,
		 balance,
		 created_at timestamp,
		)`
	_, err := s.db.Exec(str)
	return err
}

func (s *PostgresStore) CreateAccount(account *types.Account) error {
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}
func (s *PostgresStore) UpdateAccount(account *types.Account) error {
	return nil
}
func (s *PostgresStore) GetAccountByID(id int) (*types.Account, error) {
	return nil, nil
}
