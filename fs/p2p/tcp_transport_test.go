package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransport(t *testing.T) {
	listenAddr := "127.0.0.1:4396"
	tr := NewTCPTransport(listenAddr)

	assert.Equal(t, tr.listenAddress, listenAddr)

	// Server
	// tr.Start
	assert.Nil(t, tr.ListenAndAccept())
}
