package client

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/tserkov/golt/messages"
	"github.com/tserkov/golt/transport"
	"github.com/tserkov/golt/types"
)

var (
	supportedVersions = []byte{
		0x00, 0x00, 0x00, 0x01, // v1
		0x00, 0x00, 0x00, 0x00, // none
		0x00, 0x00, 0x00, 0x00, // none
		0x00, 0x00, 0x00, 0x00, // none
	}

	handshake = append(transport.Preamble, supportedVersions...)
)

type Client struct {
	cfg *Config

	conn transport.Conn
}

func New(cfg *Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &Client{cfg: cfg}, nil
}

func (c *Client) Connect() error {
	err := c.conn.Open(c.cfg.Addr, c.cfg.TLSConfig, c.cfg.ConnectionTimeout)
	if err != nil {
		return err
	}

	if err = c.handshake(); err != nil {
		return err
	}

	if err = c.init(c.cfg.Username, c.cfg.Password); err != nil {
		return err
	}

	return nil
}

// handshake performs the secret bolt club handshake (magic preamble + four supported versions).
func (c *Client) handshake() error {
	if _, err := c.conn.Write(handshake); err != nil {
		return err
	}

	v := make([]byte, 4)
	if n, err := c.conn.Read(v); err != nil {
		return err
	} else if n != 4 {
		return errors.New("Server did not respond with a valid version")
	} else if !bytes.Equal(v, supportedVersions[:4]) {
		return errors.New("Server uses an unsupported version")
	}

	return nil
}

func (c *Client) init(username, password string) error {
	init := messages.Init(
		types.String("golt/1.0"),
		types.String(username),
		types.String(password),
	)

	if err := init.Serialize(c.conn); err != nil {
		return err
	}

	msg, err := c.conn.Next()
	if err != nil {
		return err
	}

	if msg.Signature != messages.SignatureSuccess {
		return errors.New("Shit.")
	}

	fmt.Printf("%#v\n", msg)

	return nil
}

// TODO(tserkov): Connection pooling.
// Session duplicates the connection to the server.
// func (c *Client) Session() (*Conn, error) {}
