package transport

import (
	"bytes"
	"crypto/tls"
	"errors"
	"net"
	"time"

	"github.com/tserkov/golt/messages"
)

var (
	preamble = []byte{0x60, 0x60, 0xb0, 0x17}

	noSupportedVersions = []byte{0x00, 0x00, 0x00, 0x00}
	supportedVersions   = []byte{
		0x00, 0x00, 0x00, 0x01, // v1
		0x00, 0x00, 0x00, 0x00, // none
		0x00, 0x00, 0x00, 0x00, // none
		0x00, 0x00, 0x00, 0x00, // none
	}

	handshake = append(preamble, supportedVersions...)
)

type Conn struct {
	conn          net.Conn
	serverVersion []byte
	tlsConfig     *tls.Config
	uri           string
}

func (c *Conn) Connect(uri string, tlsConfig *tls.Config, timeout time.Duration) error {
	url, err := NewURL(uri)
	if err != nil {
		return err
	}

	dialer := net.Dialer{
		Timeout: timeout,
	}

	if tlsConfig != nil {
		c.conn, err = tls.DialWithDialer(&dialer, "tcp", url.Address, tlsConfig)
	} else {
		c.conn, err = dialer.Dial("tcp", url.Address)
	}
	if err != nil {
		return err
	}

	if err = c.handshake(); err != nil {
		return err
	}

	if err = c.init(); err != nil {
		return err
	}

	return nil
}

// handshake performs the secret bolt club handshake (magic preamble + four supported versions).
func (c *Conn) handshake() error {
	if _, err := c.conn.Write(handshake); err != nil {
		return err
	}

	if bytesRead, err := c.conn.Read(c.serverVersion); err != nil {
		return err
	} else if bytesRead != 4 {
		return errors.New("Server did not respond with a valid version")
	} else if bytes.Equal(c.serverVersion, noSupportedVersions) {
		return errors.New("Server uses an unsupported version")
	}

	return nil
}

func (c *Conn) init() error {
	// TODO(tserkov): provide creds
	msg := messages.Init("golt/1.0", "", "")

	if err := msg.Serialize(c.conn); err != nil {
		return err
	}

	return nil
}
