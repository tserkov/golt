package types

type Null struct{}

// Null is always encoded using the single marker byte 0xC0.
func (n Null) Encode(buf *Buffer) error {
	buf.WriteByte(MarkerNull)

	return nil
}
