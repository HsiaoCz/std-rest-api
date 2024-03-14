package store

import (
	"database/sql"

	"github.com/HsiaoCz/std-rest-api/s1mple/types"
)

type Store interface {
	CreateUser(*types.User) error
}

type MysqlStore struct {
	db *sql.DB
}

func NewMysqlStore() Store {
	return &MysqlStore{}
}

func (s *MysqlStore) init() error {
	return nil
}

func (m *MysqlStore) CreateUser(user *types.User) error {
	return nil
}
