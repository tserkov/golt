package messages

import (
	"errors"

	"github.com/tserkov/golt/types"
)

type Message = types.Structure

func Unserialize(b []byte) (*Message, error) {
	if len(b) == 0 {
		return nil, errors.New("no data to unserialize")
	}

	if b[0] == types.MarkerTinyStruct {
		return &Message{
			Signature: b[1],
		}, nil
	}

	return &Message{}, nil
}
