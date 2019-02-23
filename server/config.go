package server

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
	TLSConfig         *tls.Config
}

func (c *Config) Validate() error {
	if c.ConnectionTimeout == 0 {
		c.ConnectionTimeout = defaultConnectionTimeout
	}

	if c.Addr == "" {
		return errors.New("Missing listening address:port")
	}

	return nil
}
