package tcpc

import (
	"fmt"
	"testing"
)

func TestNewReceiver(t *testing.T) {
	receiver, err := NewReceiver[int]("127.0.0.1:3302")
	if err != nil {
		t.Error(err)
		return
	}
	msg := <-receiver.Chan
	fmt.Printf("msg: %v\n", msg)
}

func TestSender(t *testing.T) {
	_, err := NewSender[int]("127.0.0.1:3302")
	if err != nil {
		t.Error("new sender err:", err)
	}
	select {}
}
