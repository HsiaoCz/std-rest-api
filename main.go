package main

import "github.com/go-sql-driver/mysql"

func main() {
	if err := InitLogger(NewZapLoggerConf()); err != nil {
		panic(err)
	}
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "shaw123",
		Addr:                 "127.0.0.1:3306",
		DBName:               "sra",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	sqlStorage := NewMysqlStorage(cfg)
	db, err := sqlStorage.Init()
	if err != nil {
		panic(err)
	}
	store := NewStore(db)
	api := NewAPIServer("192.168.206.1:9001", store)
	api.Serve()
}
