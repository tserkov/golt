package types

import (
	"bytes"
	"testing"
)

func TestSimpleFloat(t *testing.T) {
	f := Float(1.1)

	buf := bytes.Buffer{}
	if err := f.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{0xC1, 0x3F, 0xF1, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9A}
	if !bytes.Equal(v, expected) {
		t.Errorf("expected % X, got % X", expected, v)
	}
}

func TestNegativeFloat(t *testing.T) {
	f := Float(-1.1)

	buf := bytes.Buffer{}
	if err := f.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{0xC1, 0xBF, 0xF1, 0x99, 0x99, 0x99, 0x99, 0x99, 0x9A}
	if !bytes.Equal(v, expected) {
		t.Errorf("expected % X, got % X", expected, v)
	}
}
