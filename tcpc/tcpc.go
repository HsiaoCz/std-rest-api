package tcpc

import "net"

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
	return recv, nil
}
