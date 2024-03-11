package main

import (
	"flag"
	"log"

	"github.com/HsiaoCz/std-rest-api/bank/api"
	"github.com/HsiaoCz/std-rest-api/bank/store"
)

func main() {
	listenAddr := flag.String("listenAddr", "127.0.0.1:3002", "set api server running port")
	flag.Parse()
	store, err := store.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	api := api.NewAPIServer(*listenAddr, store)
	log.Fatal(api.Run())
}
