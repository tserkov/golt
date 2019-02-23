package types

import (
	"bytes"
	"testing"
)

func TestUnboundRelationship(t *testing.T) {
	ur := UnboundRelationship(
		Integer(1),
		String("KNOWS"),
		Map{
			String("since"): Integer(1999),
		},
	)

	buf := bytes.Buffer{}
	if err := ur.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{0xB3, 0x70, 0x01, 0x85, 0x4B, 0x4E, 0x4F, 0x57, 0x53, 0xA1, 0x85, 0x73, 0x69, 0x6E, 0x63, 0x65, 0xC9, 0x07, 0xCF}
	if !bytes.Equal(v, expected) {
		t.Errorf("Expected % X, got % X", expected, v)
	}
}
