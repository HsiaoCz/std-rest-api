package main

import (
	"flag"

	"github.com/HsiaoCz/std-rest-api/gpf/gg"
)

func main() {
	var (
		JsonAddr = flag.String("JsonAddr", "127.0.0.1:9001", "listen address of service")
		GrpcAddr = flag.String("GrpcAddr", "127.0.0.1:9002", "listen address of grpc service")
	)
	flag.Parse()

	svc := gg.LoggingService{PriceServer: gg.PriceServer{}}
	go gg.MakeGRPCServer(*GrpcAddr, svc)
	jsonServer := gg.NewJSONAPIServer(*JsonAddr, svc)
	jsonServer.Run()
}
