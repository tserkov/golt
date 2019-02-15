package types

import (
	"bytes"
	"testing"
)

func TestNode(t *testing.T) {
	n := Node(
		Integer(1),
		List{},
		Map{},
	)

	buf := bytes.Buffer{}
	if err := n.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{0xB3, 0x4E, 0x01, 0x90, 0xA0}
	if !bytes.Equal(v, expected) {
		t.Errorf("Expected % X, got % X", expected, v)
	}
}
