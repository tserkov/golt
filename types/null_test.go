package types

import (
	"bytes"
	"testing"
)

func TestNull(t *testing.T) {
	n := Null{}

	b := bytes.Buffer{}
	if err := n.Encode(&b); err != nil {
		t.Error(err)
		return
	}
	v := b.Bytes()

	expected := []byte{0xC0}
	if !bytes.Equal(v, expected) {
		t.Errorf("Expected % X, got % X", v, expected)
	}
}
