package types

import (
	"bytes"
	"math"
	"testing"
)

func TestEmptyMap(t *testing.T) {
	m := Map{}

	buf := bytes.Buffer{}
	if err := m.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{MarkerTinyMap}
	if !bytes.Equal(v, expected) {
		t.Errorf("Expected % X, got % X", expected, v)
	}
}

func TestTinyMap(t *testing.T) {
	m := Map{
		String("a"): Integer(1),
	}

	buf := bytes.Buffer{}
	if err := m.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{
		byte(MarkerTinyMap + len(m)), // Map + length
		byte(MarkerTinyString + 1),   // String + length
		0x61, // String("a")
		0x01, // Integer(1)
	}
	if !bytes.Equal(v, expected) {
		t.Errorf("Expected % X, got % X", expected, v)
	}
}

func TestMap8(t *testing.T) {
	m := Map{
		String("a"): Integer(1),
		String("b"): Integer(1),
		String("c"): Integer(3),
		String("d"): Integer(4),
		String("e"): Integer(5),
		String("f"): Integer(6),
		String("g"): Integer(7),
		String("h"): Integer(8),
		String("i"): Integer(9),
		String("j"): Integer(0),
		String("k"): Integer(1),
		String("l"): Integer(2),
		String("m"): Integer(3),
		String("n"): Integer(4),
		String("o"): Integer(5),
		String("p"): Integer(6),
	}

	buf := bytes.Buffer{}
	if err := m.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	// Because golang maps aren't in any guaranteed order, it's almost impossible to test.
	// So we just check the expected length
	if v[0] != MarkerMap8 {
		t.Errorf("expected marker % X, got % X", MarkerMap8, v[0])
	}
}

func TestMap16(t *testing.T) {
	size := math.MaxUint8 + 10
	m := make(Map, size)
	for i := 0; i < size; i++ {
		m[String(i)] = Null{}
	}

	buf := bytes.Buffer{}
	if err := m.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	if v[0] != MarkerMap16 {
		t.Errorf("expected marker % X, got % X", MarkerMap16, v[0])
	}
}

func TestMap32(t *testing.T) {
	size := math.MaxUint16 * 2
	m := make(Map, size)
	for i := 0; i < size; i++ {
		m[String(string(i))] = Null{}
	}

	buf := bytes.Buffer{}
	if err := m.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	if v[0] != MarkerMap32 {
		t.Errorf("expected marker % X, got % X", MarkerMap32, v[0])
	}
}
