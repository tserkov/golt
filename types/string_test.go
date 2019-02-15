package types

import (
	"bytes"
	"math"
	"strings"
	"testing"
)

func TestEmptyString(t *testing.T) {
	b := bytes.Buffer{}
	s := String("")

	if err := s.Encode(&b); err != nil {
		t.Error(err)
		return
	}

	expected := []byte{0x80}
	v := b.Bytes()
	if !bytes.Equal(v, expected) {
		t.Errorf("expected % X, got % X", expected, v)
	}
}

func TestTinyString(t *testing.T) {
	b := bytes.Buffer{}
	s := String("a")

	if err := s.Encode(&b); err != nil {
		t.Error(err)
		return
	}

	if v := b.Bytes(); v[0] != byte(MarkerTinyString+1) {
		t.Errorf("expected marker % X, got % X", MarkerString8, v[0])
	}
}

func TestString8(t *testing.T) {
	b := bytes.Buffer{}
	s := String(strings.Repeat("a", 16))

	if err := s.Encode(&b); err != nil {
		t.Error(err)
		return
	}

	if v := b.Bytes(); v[0] != MarkerString8 {
		t.Errorf("expected marker % X, got % X", MarkerString8, v[0])
	}
}

func TestSpecialString8(t *testing.T) {
	b := bytes.Buffer{}
	s := String("En å flöt över ängen")

	if err := s.Encode(&b); err != nil {
		t.Error(err)
		return
	}

	if v := b.Bytes(); v[0] != MarkerString8 {
		t.Errorf("expected marker % X, got % X", MarkerString8, v[0])
	}
}

func TestString16(t *testing.T) {
	b := bytes.Buffer{}
	s := String(strings.Repeat("a", math.MaxUint8*2))

	if err := s.Encode(&b); err != nil {
		t.Error(err)
		return
	}

	if v := b.Bytes(); v[0] != MarkerString16 {
		t.Errorf("expected marker % X, got % X", MarkerString16, v[0])
	}
}

func TestString32(t *testing.T) {
	b := bytes.Buffer{}
	s := String(strings.Repeat("a", math.MaxUint16*2))

	if err := s.Encode(&b); err != nil {
		t.Error(err)
		return
	}

	if v := b.Bytes(); v[0] != MarkerString32 {
		t.Errorf("expected marker % X, got % X", MarkerString32, v[0])
	}
}
