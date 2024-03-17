package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/remote"
)

func main() {
	accountID := flag.Int64("account-id", 0, "id of account")
	listenAddr := flag.String("listenAddr", "127.0.0.1:9001", "listen address")

	flag.Parse()

	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}
	r := remote.New(*listenAddr, remote.Config{})
	fmt.Printf("hello %v,%v,%v\n", accountID, e, r)
}
