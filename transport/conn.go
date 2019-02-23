package transport

import (
	"crypto/tls"
	"encoding/binary"
	"net"
	"time"

	"github.com/tserkov/golt/messages"
)

var (
	Preamble = []byte{0x60, 0x60, 0xb0, 0x17}

	Version1    = []byte{0x00, 0x00, 0x00, 0x01}
	VersionNone = []byte{0x00, 0x00, 0x00, 0x00}
)

type Conn struct {
	net.Conn

	ClientName string

	tlsConfig *tls.Config
}

func (c *Conn) Open(addr string, tlsConfig *tls.Config, timeout time.Duration) error {
	dialer := net.Dialer{
		Timeout: timeout,
	}

	var err error
	if tlsConfig != nil {
		c.Conn, err = tls.DialWithDialer(&dialer, "tcp", addr, tlsConfig)
	} else {
		c.Conn, err = dialer.Dial("tcp", addr)
	}
	if err != nil {
		return err
	}

	return nil
}

func (c *Conn) Next() (*messages.Message, error) {
	lenBuf := make([]byte, 2)
	if n, err := c.Read(lenBuf); err != nil || n != 2 {
		// Read error
		c.Close()
		return nil, err
	}

	realLen := binary.BigEndian.Uint16(lenBuf)

	data := make([]byte, realLen)
	if n, err := c.Read(data); err != nil || uint16(n) != realLen {
		// Read error
		c.Close()
		return nil, err
	}

	return messages.Unserialize(data)
}

func (c *Conn) Expect(sig byte) (*messages.Message, bool, error) {
	msg, err := c.Next()
	if err != nil {
		return nil, false, err
	}

	if msg.Signature != sig {
		return nil, false, nil
	}

	return msg, true, nil
}
