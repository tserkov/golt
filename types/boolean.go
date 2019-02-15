package types

// https://boltprotocol.org/v1/#booleans
// Boolean values are encoded within a single marker byte, using 0xC3 to denote true and 0xC2 to denote false.
type Boolean bool

func (b Boolean) Encode(buf *Buffer) error {
	if b {
		buf.WriteByte(MarkerTrue)
	} else {
		buf.WriteByte(MarkerFalse)
	}

	return nil
}
