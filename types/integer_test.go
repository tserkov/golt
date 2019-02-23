package types

import (
	"bytes"
	"testing"
)

func TestSimpleInteger(t *testing.T) {
	i := Integer(1)

	buf := bytes.Buffer{}
	if err := i.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{0x01}
	if !bytes.Equal(v, expected) {
		t.Errorf("expected % X, got % X", expected, v)
	}
}

func TestInteger8Neg(t *testing.T) {
	i := Integer(-56)

	buf := bytes.Buffer{}
	if err := i.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{MarkerInt8, 0xC8}
	if !bytes.Equal(v, expected) {
		t.Errorf("expected % X, got % X", expected, v)
	}
}

func TestInteger16Neg(t *testing.T) {
	i := Integer(-16320)

	buf := bytes.Buffer{}
	if err := i.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{MarkerInt16, 0xC0, 0x40}
	if !bytes.Equal(v, expected) {
		t.Errorf("expected % X, got % X", expected, v)
	}
}

func TestInteger32(t *testing.T) {
	i := Integer(1073725440)

	buf := bytes.Buffer{}
	if err := i.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{MarkerInt32, 0x3F, 0xFF, 0xC0, 0x00}
	if !bytes.Equal(v, expected) {
		t.Errorf("expected % X, got % X", expected, v)
	}
}

func TestInteger32Neg(t *testing.T) {
	i := Integer(-1073725440)

	buf := bytes.Buffer{}
	if err := i.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{MarkerInt32, 0xC0, 0x00, 0x40, 0x00}
	if !bytes.Equal(v, expected) {
		t.Errorf("expected % X, got % X", expected, v)
	}
}

func TestMinInteger(t *testing.T) {
	i := Integer(-9223372036854775808)

	buf := bytes.Buffer{}
	if err := i.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{MarkerInt64, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	if !bytes.Equal(v, expected) {
		t.Errorf("expected % X, got % X", expected, v)
	}
}

func TestMaxInteger(t *testing.T) {
	i := Integer(9223372036854775807)

	buf := bytes.Buffer{}
	if err := i.Encode(&buf); err != nil {
		t.Error(err)
		return
	}
	v := buf.Bytes()

	expected := []byte{MarkerInt64, 0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	if !bytes.Equal(v, expected) {
		t.Errorf("expected % X, got % X", expected, v)
	}
}
