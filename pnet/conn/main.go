package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:30001")
	if err != nil {
		log.Fatal(err)
	}
	// listener
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go handleTcpAccept(conn)
	}
}

func handleTcpAccept(conn *net.TCPConn) {
	fmt.Println(conn.RemoteAddr().String())
}
