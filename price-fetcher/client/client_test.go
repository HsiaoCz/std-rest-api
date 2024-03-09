package client

import (
	"context"
	"log"
	"testing"
)

func TestClient(t *testing.T) {
	client := New("http://127.0.0.1:3301")
	pirce, err := client.FetchPrice(context.Background(), "GG")
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("the price response: %+v\n", pirce)
}
