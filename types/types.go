package types

import (
	"bytes"
)

type Buffer = bytes.Buffer

type Any interface {
	Encode(*Buffer) error
}
