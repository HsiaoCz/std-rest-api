package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/HsiaoCz/std-rest-api/codebeast/structure/api"
	"github.com/HsiaoCz/std-rest-api/codebeast/structure/storage"
)

func main() {
	listenAddr := flag.String("listenAddr", ":30001", "set the server address")
	flag.Parse()

	storage := storage.NewMongoStorage()

	server := api.NewServer(*listenAddr, storage)
	fmt.Println("server is running on port:", *listenAddr)
	log.Fatal(server.Start())
}
