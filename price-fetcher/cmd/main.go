package main

import (
	"flag"

	"github.com/HsiaoCz/std-rest-api/price-fetcher/gg"
)

func main() {
	listenAddr := flag.String("listenAddr", "127.0.0.1:3301", "set the server listen address")
	flag.Parse()
	svc := gg.NewLoggingService(&gg.Pricefetch{})

	server := gg.NewJSONAPIServer(*listenAddr, svc)
	
	server.Run()

}
