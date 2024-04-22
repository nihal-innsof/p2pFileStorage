package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddresss := ":4000"
	tr := NewTCPTransport(listenAddresss)

	assert.Equal(t, tr.listenAddress, listenAddresss)
}
