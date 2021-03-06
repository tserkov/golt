package client

import (
	"crypto/tls"
	"errors"
	"time"
)

var (
	defaultConnectionTimeout time.Duration = 60 * time.Second
)

type Config struct {
	Addr              string
	ConnectionTimeout time.Duration
	Password          string
	TLSConfig         *tls.Config
	Username          string
}

func (c *Config) Validate() error {
	if c.ConnectionTimeout == 0 {
		c.ConnectionTimeout = defaultConnectionTimeout
	}

	if c.Addr == "" {
		return errors.New("Missing server address:port")
	}

	if c.Username == "" && c.Password != "" {
		return errors.New("Username must be set when a password is set")
	}

	return nil
}
