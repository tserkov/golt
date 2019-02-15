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
	ConnectionTimeout time.Duration
	URI               string
	TLSConfig         *tls.Config
}

func (c *Config) Validate() error {
	if c.ConnectionTimeout == 0 {
		c.ConnectionTimeout = defaultConnectionTimeout
	}

	if c.URI == "" {
		return errors.New("Missing server URI")
	}

	return nil
}
