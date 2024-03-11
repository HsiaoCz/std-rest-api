package tcpc

import (
	"encoding/gob"
	"log"
	"net"
	"time"
)

// func NewChannel[T any]() (*Sender[T], *Reciver[T]) {}

type TCPC[T any] struct {
	listenAddr string
	remoteAddr string

	SendChan chan T
	RecvChan chan T

	outboundConn net.Conn
	ln           net.Listener
}

func New[T any](listenAddr, remoteAddr string) (*TCPC[T], error) {
	tcpc := &TCPC[T]{
		listenAddr: listenAddr,
		remoteAddr: remoteAddr,
		SendChan:   make(chan T, 10),
		RecvChan:   make(chan T, 10),
	}
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, err
	}
	tcpc.ln = ln
	go tcpc.acceptLoop()
	return tcpc, nil
}

func (t *TCPC[T]) dialRemote() {
	conn, err := net.Dial("tcp", t.remoteAddr)
	if err != nil {
		log.Printf("dial error (%s)", err)
		time.Sleep(time.Second)
		t.dialRemote()
	}
	t.outboundConn = conn
}

func (t *TCPC[T]) acceptLoop() {
	defer func() {
		t.ln.Close()
	}()
	for {
		conn, err := t.ln.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			return
		}
		log.Printf("sender connected %s", conn.RemoteAddr())
		go t.handleConn(conn)
	}
}

func (t *TCPC[T]) handleConn(conn net.Conn) {
	for {
		var msg T
		if err := gob.NewDecoder(conn).Decode(&msg); err != nil {
			log.Printf("Decode the msg err: %v\n", err)
			continue
		}
		t.RecvChan <- msg
	}
}

var defaultDialInterval = 3 * time.Second

type Sender[T any] struct {
	Chan chan T

	remoteAddr   string
	outboundConn net.Conn

	dialInterval time.Duration
}

func NewSender[T any](remoteAddr string) (*Sender[T], error) {
	sender := &Sender[T]{
		Chan:         make(chan T),
		remoteAddr:   remoteAddr,
		dialInterval: defaultDialInterval,
	}
	go sender.dialRemote()
	go sender.loop()
	return sender, nil
}

func (s *Sender[T]) dialRemote() {
	conn, err := net.Dial("tcp", s.remoteAddr)
	if err != nil {
		log.Printf("dial error (%s)  retrying in (%v) \n", err, s.dialInterval)
		time.Sleep(s.dialInterval)
		s.dialRemote()
	}
	s.outboundConn = conn
}

func (s *Sender[T]) loop() {
	for {
		msg := <-s.Chan
		log.Println("sending msg over the write:", msg)
		if err := gob.NewEncoder(s.outboundConn).Encode(msg); err != nil {
			log.Printf("gob encode the msg err: %v\n", err)
			continue
		}
	}
}

type Reciver[T any] struct {
	Chan       chan T
	listenAddr string
	listener   net.Listener
}

func NewReceiver[T any](listenAddr string) (*Reciver[T], error) {
	recv := &Reciver[T]{
		Chan:       make(chan T),
		listenAddr: listenAddr,
	}
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, err
	}
	recv.listener = ln

	go recv.acceptLoop()

	return recv, nil
}

func (r *Reciver[T]) acceptLoop() {
	defer func() {
		r.listener.Close()
	}()
	for {
		conn, err := r.listener.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			return
		}
		log.Printf("sender connected %s", conn.RemoteAddr())
		go r.handleConn(conn)
	}
}

func (r *Reciver[T]) handleConn(conn net.Conn) {
	for {
		var msg T
		if err := gob.NewDecoder(conn).Decode(&msg); err != nil {
			log.Printf("Decode the msg err: %v\n", err)
			continue
		}
		r.Chan <- msg
	}
}
