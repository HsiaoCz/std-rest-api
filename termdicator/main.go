package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/VictorLowther/btree"
	"github.com/gorilla/websocket"
)

const (
	wsendpoint = "wss://fstream.binance.com/stream?streams=btcusdt@depth"
)

func byBestBid(a, b *OrderBookEntry) bool {
	return a.Price > b.Price
}

func byBestAsk(a, b *OrderBookEntry) bool {
	return a.Price < b.Price
}

type OrderBookEntry struct {
	Price  float64
	Volume float64
}

type Orderbook struct {
	Asks *btree.Tree[*OrderBookEntry]
	Bids *btree.Tree[*OrderBookEntry]
}

func NewOrderBook() *Orderbook {
	return &Orderbook{
		Asks: btree.New(byBestAsk),
		Bids: btree.New(byBestBid),
	}
}

func (ob *Orderbook) handleDepthResponse(res BinanceDepthResult) {
	for _, ask := range res.Asks {
		price, _ := strconv.ParseFloat(ask[0], 64)
		volume, _ := strconv.ParseFloat(ask[1], 64)
		entry := &OrderBookEntry{
			Price:  price,
			Volume: volume,
		}
		fmt.Printf("%+v\n", entry)
	}
}

type BinanceDepthResult struct {
	// price | size (volume)
	Asks [][]string `json:"a"`
	Bids [][]string `json:"b"`
}

type BinanceDepthResponse struct {
	Stream string             `json:"stream"`
	Data   BinanceDepthResult `json:"data"`
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(wsendpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	var (
		ob     = NewOrderBook()
		result BinanceDepthResponse
	)
	for {
		if err := conn.ReadJSON(&result); err != nil {
			log.Fatal(err)
		}
		ob.handleDepthResponse(result.Data)
	}
}
