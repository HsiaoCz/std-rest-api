package main

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/HsiaoCz/std-rest-api/gserver/types"
	"github.com/gorilla/websocket"
)

const wsServerEndpoint = "ws://localhost:40000/ws"

type GameClient struct {
	conn     *websocket.Conn
	clientID int
	username string
}

func newGameClient(conn *websocket.Conn, username string) *GameClient {
	return &GameClient{
		clientID: rand.New(rand.NewSource(time.Now().UnixNano())).Intn(math.MaxInt),
		username: username,
		conn:     conn,
	}
}

func (g *GameClient) login() error {
	b, err := json.Marshal(types.Login{
		ClientID: g.clientID,
		Username: g.username,
	})
	if err != nil {
		return err
	}
	msg := types.WsMessage{
		Type: "login",
		Data: b,
	}
	return g.conn.WriteJSON(&msg)
}

func main() {
	dialer := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, _, err := dialer.Dial(wsServerEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := newGameClient(conn, "jhonason")
	if err := client.login(); err != nil {
		log.Fatal(err)
	}
}
