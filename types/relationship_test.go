package types

import (
	"bytes"
	"testing"
)

func TestRelationship(t *testing.T) {
	r := Relationship(
		Integer(1),
		Integer(2),
		Integer(3),
		String("a"),
		Map{},
	)

	buf := bytes.Buffer{}
	if err := r.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{0xB5, 0x52, 0x01, 0x02, 0x03, 0x81, 0x61, 0xA0}
	if !bytes.Equal(v, expected) {
		t.Errorf("Expected % X, got % X", expected, v)
	}
}
