package types

import (
	"encoding/binary"
	"errors"
	"math"
)

// https://boltprotocol.org/v1/#strings
// String data is represented as UTF-8 encoded binary data.
// Note that sizes used for string are the byte counts of the UTF-8 encoded data, not the character count of the original string.
type String string

func (s String) Encode(buf *Buffer) error {
	l := len(s)
	if l <= 15 {
		buf.WriteByte(byte(MarkerTinyString + l))
	} else if l <= math.MaxUint8 {
		buf.WriteByte(MarkerString8)
		binary.Write(buf, binary.BigEndian, int8(l))
	} else if l <= math.MaxUint16 {
		buf.WriteByte(MarkerString16)
		binary.Write(buf, binary.BigEndian, int16(l))
	} else if l <= math.MaxUint32 {
		buf.WriteByte(MarkerString32)
		binary.Write(buf, binary.BigEndian, int32(l))
	} else {
		return errors.New("String is too long.")
	}

	buf.WriteString(string(s))

	return nil
}
