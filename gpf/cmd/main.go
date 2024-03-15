package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/HsiaoCz/std-rest-api/gpf/client"
	"github.com/HsiaoCz/std-rest-api/gpf/gg"
	"github.com/HsiaoCz/std-rest-api/gpf/proto"
)

func main() {
	var (
		JsonAddr = flag.String("JsonAddr", "127.0.0.1:9001", "listen address of service")
		GrpcAddr = flag.String("GrpcAddr", "127.0.0.1:9002", "listen address of grpc service")
		ctx      = context.Background()
	)
	flag.Parse()

	svc := gg.LoggingService{PriceServer: gg.PriceServer{}}

	grpcClient, err := client.NewGRPCClient("127.0.0.1:9002")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		time.Sleep(time.Second * 3)
		resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "GG"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", resp)
	}()

	go gg.MakeGRPCServer(*GrpcAddr, svc)
	jsonServer := gg.NewJSONAPIServer(*JsonAddr, svc)
	jsonServer.Run()
}
