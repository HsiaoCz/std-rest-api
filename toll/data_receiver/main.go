package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HsiaoCz/std-rest-api/toll/types"
	"github.com/gorilla/websocket"
)

func main() {
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
