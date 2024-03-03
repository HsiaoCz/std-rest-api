package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HsiaoCz/std-rest-api/toll/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gorilla/websocket"
)

// kafka topic

const kafkaTopic = "obudata"

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := kafkaTopic
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("test producting"),
	}, nil)
	return
	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
	recv := NewDataRecevier()
	http.HandleFunc("/ws", recv.wsHandler)
	http.ListenAndServe("127.0.0.1:3001", nil)
}

type DataReceiver struct {
	msgchan chan types.OBUData
	conn    *websocket.Conn
}

func (dr *DataReceiver) wsHandler(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn

	go dr.wsReceiveLoop()
}

// the channel always block when is full
func NewDataRecevier() *DataReceiver {
	return &DataReceiver{
		msgchan: make(chan types.OBUData, 128),
	}
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("new obu connected client connected")
	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error:", err)
			continue
		}
		fmt.Printf("received OBU data from [%d] :: <lat %.2f, long %.2f>  \n", data.OBUID, data.Lat, data.Long)
		dr.msgchan <- data
	}
}
