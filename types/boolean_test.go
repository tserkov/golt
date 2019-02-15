package types

import (
	"bytes"
	"testing"
)

func TestTrue(t *testing.T) {
	b := Boolean(true)

	buf := bytes.Buffer{}
	if err := b.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{0xC3}
	if !bytes.Equal(v, expected) {
		t.Errorf("Expected % X, got % X", v, expected)
	}
}

func TestFalse(t *testing.T) {
	b := Boolean(false)

	buf := bytes.Buffer{}
	if err := b.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{0xC2}
	if !bytes.Equal(v, expected) {
		t.Errorf("Expected % X, got % X", v, expected)
	}
}
