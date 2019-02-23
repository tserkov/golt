package types

import (
	"bufio"
	"bytes"
	"math"
	"testing"
)

func TestEmptyStructure(t *testing.T) {
	b := bytes.Buffer{}
	s := Structure{
		Signature: 0x01,
	}

	if err := s.Encode(&b); err != nil {
		t.Error(err)
		return
	}

	if v := b.Bytes(); v[0] != MarkerTinyStruct {
		t.Errorf("expected marker % X, got % X", MarkerTinyStruct, v[0])
	}
}

func TestTinyStructure(t *testing.T) {
	b := bytes.Buffer{}
	s := Structure{
		Fields: []Value{
			Integer(1),
			Integer(2),
			Integer(3),
		},
		Signature: 0x01,
	}

	if err := s.Encode(&b); err != nil {
		t.Error(err)
		return
	}

	if v := b.Bytes(); v[0] != byte(MarkerTinyStruct+3) {
		t.Errorf("expected marker % X, got % X", byte(MarkerTinyStruct+3), v[0])
	}
}

func TestStructure8(t *testing.T) {
	b := bytes.Buffer{}
	s := Structure{
		Fields: []Value{
			Integer(1),
			Integer(2),
			Integer(3),
			Integer(4),
			Integer(5),
			Integer(6),
			Integer(7),
			Integer(8),
			Integer(9),
			Integer(0),
			Integer(1),
			Integer(2),
			Integer(3),
			Integer(4),
			Integer(5),
			Integer(6),
		},
		Signature: 0x01,
	}

	if err := s.Encode(&b); err != nil {
		t.Error(err)
		return
	}

	if v := b.Bytes(); v[0] != MarkerStruct8 {
		t.Errorf("expected marker % X, got % X", MarkerStruct8, v[0])
	}
}

func TestStructure16(t *testing.T) {
	fields := make([]Value, math.MaxUint8*2)
	for i := 0; i < len(fields); i++ {
		fields[i] = Null{}
	}

	b := bytes.Buffer{}
	s := Structure{
		Fields:    fields,
		Signature: 0x01,
	}

	if err := s.Encode(&b); err != nil {
		t.Error(err)
		return
	}

	if v := b.Bytes(); v[0] != MarkerStruct16 {
		t.Errorf("expected marker %X, got %X", MarkerStruct16, v[0])
	}
}

func TestStructureSerialize(t *testing.T) {
	s := Structure{
		Fields: []Value{
			Integer(1),
			String("a"),
			Boolean(false),
		},
		Signature: 0x01,
	}

	buf := bytes.Buffer{}
	w := bufio.NewWriter(&buf)

	if err := s.Serialize(w); err != nil {
		t.Error(err)
		return
	}

	w.Flush()

	expected := []byte{0x00, 0x06, 0xB3, 0x01, 0x01, 0x81, 0x61, 0xC2, 0x00, 0x00}
	if v := buf.Bytes(); !bytes.Equal(v, expected) {
		t.Errorf("expected % X, got % X", expected, v)
	}
}

func BenchmarkStructureSerialize(b *testing.B) {
	s := Structure{
		Fields: []Value{
			Integer(1),
			String("a"),
			Boolean(false),
		},
		Signature: 0x01,
	}

	var buf bytes.Buffer
	w := bufio.NewWriterSize(&buf, b.N*10) // The above struct serializes to 10 bytes.

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := s.Serialize(w); err != nil {
			b.Fatal(err)
		}
	}
}
