package main

import (
	"context"
	"fmt"
	"log"

	"github.com/HsiaoCz/std-rest-api/price-fetcher/gg"
)

func main() {
	svc := gg.NewLoggingService(&gg.Pricefetch{})
	price, err := svc.FetchPrice(context.Background(), "GG")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("price : %v\n", price)
}
