package transport

import (
	"crypto/tls"
	"net"
)

type Listener struct {
	net.Listener
}

func (l *Listener) Accept(addr string, _ *tls.Config, fn func(Conn)) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	l.Listener = ln

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go fn(Conn{Conn: conn})
	}
}
