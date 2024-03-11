package main

import (
	"log"

	"github.com/HsiaoCz/std-rest-api/tcpc"
)

// type Server[T any] struct {
// 	receiver *tcpc.Reciver[T]
// }

// func (s *Server[T]) recvLoop() {
// 	for {
// 		<-s.receiver.Chan
// 	}
// }

// func main() {
// 	// local
// 	receiver, err := tcpc.NewReceiver[int](":3301")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// remote
// 	sender, err := tcpc.NewSender[int](":3301")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// option
// 	// tcpc := New(":3001")
// 	// msg:= <- tcpc.Recv
// 	// tcpc.Send <- msg

// 	time.Sleep(time.Second * 2)
// 	sender.Chan <- 100
// 	msg := <-receiver.Chan
// 	fmt.Println("receive from  channel over the wire:", msg)
// }

func main() {
	channel, err := tcpc.New[string](":30001", ":40001")
	if err != nil {
		log.Fatal(err)
	}

	channel.RecvChan <- "GG"
}
