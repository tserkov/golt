package server

import (
	"bytes"
	"net"

	"github.com/tserkov/golt/messages"
	"github.com/tserkov/golt/transport"
	"github.com/tserkov/golt/types"
)

type Server struct {
	cfg *Config

	ln transport.Listener
}

func New() *Server {
	return &Server{}
}

func (s *Server) Addr() net.Addr {
	return s.ln.Addr()
}

func (s *Server) Listen(addr string) error {
	err := s.ln.Accept(s.cfg.Addr, s.cfg.TLSConfig, s.accept)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) accept(conn transport.Conn) {
	preambleBuf := make([]byte, 4)
	if n, err := conn.Read(preambleBuf); err != nil || n != 4 {
		// Read error
		conn.Close()
		return
	}

	// Check bolt magic preamble.
	if !bytes.Equal(preambleBuf, transport.Preamble) {
		conn.Close()
		return
	}

	// Read client's 4 potential supported versions.
	versBuf := make([]byte, 16)
	if n, err := conn.Read(versBuf); err != nil || n != 16 {
		// Read error
		conn.Close()
		return
	}

	// Of course, we only support version 1.
	for i := 0; i < 15; i += 4 {
		if !bytes.Equal(versBuf[i:4], transport.Version1) {
			conn.Write(transport.VersionNone)
			conn.Close()
			return
		}
	}

	// Confirm our supported version.
	conn.Write(transport.Version1)

	initMsg, err := s.awaitInit(conn)
	if err != nil {
		return
	}

	// TODO(tserkov): Authentication.

	conn.ClientName = string(initMsg.Fields[0].(types.String))

	// SUCCESS
	messages.Success(types.Map{}).Serialize(conn)
}

func (s *Server) awaitInit(conn transport.Conn) (msg *messages.Message, err error) {
	ok := false
	for !ok {
		msg, ok, err = conn.Expect(messages.SignatureInit)
		if err != nil {
			return nil, err
		}

		if !ok {
			if err = s.mustAckFailure(conn); err != nil {
				return nil, err
			}
		}
	}

	return msg, nil
}

func (s *Server) mustAckFailure(conn transport.Conn) error {
	// Send FAILURE
	if err := messages.Failure(types.Map{}).Serialize(conn); err != nil {
		return err
	}

	var err error

	// Await ACK_FAILURE
	ok := false
	for !ok {
		_, ok, err = conn.Expect(messages.SignatureAckFailure)
		if err != nil {
			return err
		}
	}

	return nil
}
