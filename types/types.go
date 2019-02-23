package types

import (
	"bytes"
)

type Buffer = bytes.Buffer

type Value interface {
	Encode(*Buffer) error
}
