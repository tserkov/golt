package transport

import (
	"net"
	"os"
	"sync"
	"testing"
	"time"
)

var testServerAddr net.Addr
var connectionCount uint8
var mu sync.RWMutex

func TestMain(m *testing.M) {
	s := Listener{}

	ch := make(chan error, 1)
	go func() {
		ch <- s.Accept("127.0.0.1:0", nil, func(c Conn) {
			mu.Lock()
			connectionCount += 1
			mu.Unlock()

			c.Write([]byte{
				0x00, // len
				0x02, // len
				0xB0, // TINY_STRUCT (0-field)
				0xff, // signature (fake)
			})
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

func TestConnect(t *testing.T) {
	var c Conn
	if err := c.Open(testServerAddr.String(), nil, 2*time.Second); err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	mu.RLock()
	defer mu.RUnlock()
	if connectionCount == 0 {
		t.Fatal("no connection registered")
	}
}

func TestNext(t *testing.T) {
	var c Conn
	if err := c.Open(testServerAddr.String(), nil, 2*time.Second); err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	msg, err := c.Next()
	if err != nil {
		t.Fatal(err)
		return
	}

	if msg.Signature != 0xff {
		t.Fatal("bad message")
	}
}

func TestExpectOK(t *testing.T) {
	var c Conn
	if err := c.Open(testServerAddr.String(), nil, 2*time.Second); err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	_, ok, err := c.Expect(0xff)
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		t.Fatal("did not get expected signature")
	}
}

func TestExpectBad(t *testing.T) {
	var c Conn
	if err := c.Open(testServerAddr.String(), nil, 2*time.Second); err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	_, ok, err := c.Expect(0x01)
	if err != nil {
		t.Fatal(err)
	}

	if ok {
		t.Fatal("did not get expected signature")
	}
}
