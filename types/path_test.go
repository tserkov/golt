package types

import (
	"bytes"
	"testing"
)

func TestPath(t *testing.T) {
	p := Path(
		List{},
		List{},
		List{},
	)

	buf := bytes.Buffer{}
	if err := p.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{0xB3, 0x50, 0x90, 0x90, 0x90}
	if !bytes.Equal(v, expected) {
		t.Errorf("Expected % X, got % X", expected, v)
	}
}
