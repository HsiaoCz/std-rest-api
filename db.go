package main

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type MysqlStorage struct {
	db *sql.DB
}

func NewMysqlStorage(cfg mysql.Config) *MysqlStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	zap.L().Info("connect db successful")
	return &MysqlStorage{db: db}
}

func (s *MysqlStorage) Init() (*sql.DB, error) {
	// init the tables
	return s.db, nil
}
