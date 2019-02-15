package types

import (
	"encoding/binary"
	"errors"
	"math"
)

// Maps are sized sequences of pairs of keys and values and permit a mixture of types within the same map.
// The size of a map denotes the number of pairs within that map, not the total packed byte size.
// Keys are unique within a map, however the serialization format notably technically allows duplicate keys to be sent.
// Though if duplicate keys are sent, this is a violation of the bolt protocol and an error will occur.
// In Go, the order of the map is not guaranteed.

type Map map[String]Any

func (m Map) Encode(buf *Buffer) error {
	l := len(m)

	// Set marker + length.
	if l <= 15 {
		buf.WriteByte(byte(MarkerTinyMap + l))
	} else if l <= math.MaxUint8 {
		buf.WriteByte(MarkerMap8)
		binary.Write(buf, binary.BigEndian, int8(l))
	} else if l <= math.MaxUint16 {
		buf.WriteByte(MarkerMap16)
		binary.Write(buf, binary.BigEndian, int16(l))
	} else if l <= math.MaxUint32 {
		buf.WriteByte(MarkerMap32)
		binary.Write(buf, binary.BigEndian, int32(l))
	} else {
		return errors.New("Map is too long")
	}

	for k, v := range m {
		if err := k.Encode(buf); err != nil {
			return err
		}

		if err := v.Encode(buf); err != nil {
			return err
		}
	}

	return nil
}
