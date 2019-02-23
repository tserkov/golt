package types

import (
	"encoding/binary"
)

// https://boltprotocol.org/v1/#floats
// These are double-precision floating points for approximations of any number, notably for
// representing fractions and decimal numbers. Floats are encoded as a single 0xC1 marker byte
// followed by 8 bytes, formatted according to the IEEE 754 floating-point "double format" bit layout.
type Float float64

func (f Float) Encode(buf *Buffer) error {
	buf.WriteByte(MarkerFloat64)

	if err := binary.Write(buf, binary.BigEndian, f); err != nil {
		return err
	}

	return nil
}
