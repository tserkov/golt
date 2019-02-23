package types

import (
	"encoding/binary"
	"errors"
	"math"
)

// https://boltprotocol.org/v1/#lists
// Lists are heterogeneous sequences of values and permit a mixture of types within the same list.
// The size of a list denotes the number of items within that list, not the total packed byte size.
type List []Value

func (l List) Encode(buf *Buffer) error {
	length := len(l)

	// Set marker + length.
	if length <= 15 {
		buf.WriteByte(byte(MarkerTinyList + length))
	} else if length <= math.MaxUint8 {
		buf.WriteByte(MarkerList8)
		binary.Write(buf, binary.BigEndian, int8(length))
	} else if length <= math.MaxUint16 {
		buf.WriteByte(MarkerList16)
		binary.Write(buf, binary.BigEndian, int16(length))
	} else if length <= math.MaxUint32 {
		buf.WriteByte(MarkerList32)
		binary.Write(buf, binary.BigEndian, int32(length))
	} else {
		return errors.New("List is too long")
	}

	for _, v := range l {
		if err := v.Encode(buf); err != nil {
			return err
		}
	}

	return nil
}
