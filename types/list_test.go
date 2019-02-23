package types

import (
	"bytes"
	"math"
	"testing"
)

func TestEmptyList(t *testing.T) {
	l := List{}

	buf := bytes.Buffer{}
	if err := l.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{MarkerTinyList}
	if !bytes.Equal(v, expected) {
		t.Errorf("Expected % X, got % X", expected, v)
	}
}

func TestTinyList(t *testing.T) {
	l := List{
		Integer(1),
		Integer(2),
		Integer(3),
	}

	buf := bytes.Buffer{}
	if err := l.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{byte(MarkerTinyList + len(l)), 0x01, 0x02, 0x03}
	if !bytes.Equal(v, expected) {
		t.Errorf("Expected % X, got % X", expected, v)
	}
}

func TestList8(t *testing.T) {
	l := List{
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
		Integer(7),
		Integer(8),
		Integer(9),
		Integer(0),
	}

	buf := bytes.Buffer{}
	if err := l.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{MarkerList8, 0x14, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x00}
	if !bytes.Equal(v, expected) {
		t.Errorf("Expected % X, got % X", expected, v)
	}
}

func TestList16(t *testing.T) {
	l := make(List, math.MaxUint16)
	for i := 0; i < math.MaxUint16; i++ {
		l[i] = Null{}
	}

	buf := bytes.Buffer{}
	if err := l.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{MarkerList16, 0xFF, 0xFF}
	if !bytes.Equal(v[:3], expected) || len(v) != math.MaxUint16+len(expected) {
		t.Errorf("Expected % X, got % X", expected, v)
	}
}

func TestList32(t *testing.T) {
	l := make(List, math.MaxUint16+10)
	for i := 0; i < math.MaxUint16+10; i++ {
		l[i] = Null{}
	}

	buf := bytes.Buffer{}
	if err := l.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{MarkerList32, 0x00, 0x01, 0x00, 0x09}
	if !bytes.Equal(v[:5], expected) || len(v) != math.MaxUint16+10+len(expected) {
		t.Errorf("Expected % X, got % X", expected, v)
	}
}
