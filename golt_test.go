package golt

import (
	"testing"
)

func TestNew(t *testing.T) {
	_, err := New(&Config{
		URI: "bolt://localhost",
	})

	if err != nil {
		t.Error(err)
	}
}
