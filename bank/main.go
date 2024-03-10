package main

import (
	"flag"
	"log"

	"github.com/HsiaoCz/std-rest-api/bank/api"
)

func main() {
	listenAddr := flag.String("listenAddr", "127.0.0.1:3002", "set api server running port")
	flag.Parse()
	api := api.NewAPIServer(*listenAddr)
	log.Fatal(api.Run())
}
