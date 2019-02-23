package client

import (
	"net"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/tserkov/golt/transport"
)

var testServerAddr net.Addr
var connectionCount uint8
var mu sync.RWMutex

func TestMain(m *testing.M) {
	s := transport.Listener{}

	ch := make(chan error, 1)
	go func() {
		ch <- s.Accept("127.0.0.1:0", nil, func(transport.Conn) {
			connectionCount += 1
		})
	}()

	// Give the server 1 second to start or error out.
	select {
	case err := <-ch:
		panic(err)
	case <-time.After(1 * time.Second):
	}

	testServerAddr = s.Addr()

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	_, err := New(&Config{
		Addr: "localhost:7687",
	})

	if err != nil {
		t.Error(err)
	}
}
