package types

import (
	"encoding/binary"
	"errors"
	"math"
)

// https://boltprotocol.org/v1/#ints
// Integer values occupy either 1, 2, 3, 5 or 9 bytes depending on magnitude and are stored as big-endian signed values.
// Several markers are designated specifically as TINY_INT values and can therefore be used to pass a small number in a single byte.
// These markers can be identified by a zero high-order bit or by a high-order nibble containing only ones.
type Integer int64

func (i Integer) Encode(buf *Buffer) error {
	if i >= math.MinInt64 && i < math.MinInt32 { // Write as INT_64
		buf.WriteByte(MarkerInt64)
		binary.Write(buf, binary.BigEndian, i)

	} else if i >= math.MinInt32 && i < math.MinInt16 { // Write as INT_32
		buf.WriteByte(MarkerInt32)
		binary.Write(buf, binary.BigEndian, int32(i))

	} else if i >= math.MinInt16 && i < math.MinInt8 { // Write as INT_16
		buf.WriteByte(MarkerInt16)
		binary.Write(buf, binary.BigEndian, int16(i))

	} else if i >= math.MinInt8 && i < -16 { // Write as INT_8
		buf.WriteByte(MarkerInt8)
		binary.Write(buf, binary.BigEndian, int8(i))

	} else if i >= -16 && i <= math.MaxInt8 { // Write as TINY_INT
		buf.WriteByte(byte(int8(MarkerTinyInt + i)))

	} else if i > math.MaxInt8 && i <= math.MaxInt16 { // Write as INT_16
		buf.WriteByte(MarkerInt16)
		binary.Write(buf, binary.BigEndian, int16(i))

	} else if i > math.MaxInt16 && i <= math.MaxInt32 { // Write as INT_32
		buf.WriteByte(MarkerInt32)
		binary.Write(buf, binary.BigEndian, int32(i))

	} else if i > math.MaxInt32 && i <= math.MaxInt64 { // Write as INT_64
		buf.WriteByte(MarkerInt64)
		binary.Write(buf, binary.BigEndian, i)

	} else {
		return errors.New("Integer is too big")
	}

	return nil
}
