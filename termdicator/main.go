package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

const (
	wsendpoint = "wss://fstream.binance.com/stream?streams=btcusdt@depth"
)

type BinanceOrderbookResult struct {
	Asks [][]string `json:"a"`
	Bids [][]string `json:"b"`
}

type BinanceDepthResponse struct {
	Stream string                 `json:"stream"`
	Data   BinanceOrderbookResult `json:"data"`
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(wsendpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	var result BinanceDepthResponse
	for {
		if err := conn.ReadJSON(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}
