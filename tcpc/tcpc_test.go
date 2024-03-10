package tcpc

import (
	"fmt"
	"testing"
)

func TestNewReceiver(t *testing.T) {
	receiver, err := NewReceiver[string]("127.0.0.1:3302")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("receiver: %+v\n", receiver)
}
