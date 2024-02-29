package main

import (
	"fmt"
	"log"

	"github.com/HsiaoCz/std-rest-api/fs/p2p"
)

func main() {
	fmt.Println("We Gucci")
	tr := p2p.NewTCPTransport("127.0.0.1:4396")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
