package types

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math"
)

// https://boltprotocol.org/v1/#structures
// Structures represent composite values and consist, beyond the marker, of a single byte signature
// followed by a sequence of fields, each an individual value. The size of a structure is measured
// as the number of fields, not the total packed byte size.
type Structure struct {
	Signature byte
	Fields    []Any
}

func (s Structure) Encode(buf *Buffer) error {
	l := len(s.Fields)
	if l <= 15 {
		buf.WriteByte(byte(MarkerTinyStruct + l))
	} else if l <= math.MaxUint8 {
		buf.WriteByte(MarkerStruct8)
		binary.Write(buf, binary.BigEndian, int8(l))
	} else if l <= math.MaxUint16 {
		buf.WriteByte(MarkerStruct16)
		binary.Write(buf, binary.BigEndian, int16(l))
	} else {
		return errors.New("Structure is too large.")
	}

	buf.WriteByte(s.Signature)

	for _, f := range s.Fields {
		if err := f.Encode(buf); err != nil {
			return err
		}
	}

	return nil
}

func (s *Structure) Serialize(w io.Writer) error {
	var buf bytes.Buffer
	if err := s.Encode(&buf); err != nil {
		return err
	}

	l := buf.Len()

	// Length
	if _, err := w.Write([]byte{uint8(l >> 8), uint8(l & 0xff)}); err != nil {
		return err
	}

	// Payload
	if _, err := buf.WriteTo(w); err != nil {
		return err
	}

	// Tail
	if _, err := w.Write(MarkerEnd); err != nil {
		return err
	}

	return nil
}
