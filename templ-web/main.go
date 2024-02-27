package main

import (
	"fmt"
	"log/slog"

	"github.com/HsiaoCz/std-rest-api/templ-web/handlers"
	"github.com/HsiaoCz/std-rest-api/templ-web/log"
	"github.com/HsiaoCz/std-rest-api/templ-web/store"
	"gorm.io/gorm"
)

var (
	addr = "127.0.0.1:9001"
)

func main() {
	if err := log.InitLogger(log.NewZapLoggerConf()); err != nil {
		slog.Error("init logger err:", err)
		return
	}
	handler := handlers.NewHandler(addr, store.NewStorage(&gorm.DB{}), handlers.NewUserHandler(store.NewStorage(&gorm.DB{})))
	if err := handler.Serve(); err != nil {
		slog.Error("handler serve err:", err)
		return
	}
	fmt.Println("templ and web components")
}
