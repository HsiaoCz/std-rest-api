package main

import (
	"flag"

	"github.com/HsiaoCz/std-rest-api/gpf/gg"
)

func main() {
	listenAddr := flag.String("listenAddr", "127.0.0.1:9001", "listen address of service")
	flag.Parse()

	svc := gg.LoggingService{PriceServer: gg.PriceServer{}}
	server := gg.NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
