package main

import "database/sql"

type Store interface {
	// users
	CreateUser(*User) error
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateUser(*User) error {
	return nil
}
