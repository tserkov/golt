package messages

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/tserkov/golt/types"
)

func TestIgnored(t *testing.T) {
	msg := Ignored()

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	if err := msg.Serialize(w); err != nil {
		t.Error(err)
		return
	}
	w.Flush()

	v := buf.Bytes()

	// Make sure the message is a 0-field struct
	if v[2] != types.MarkerTinyStruct {
		t.Errorf("expected message type % X, got % X", types.MarkerTinyStruct, v[2])
	}

	// Make sure the message signature is IGNORED
	if v[3] != SignatureIgnored {
		t.Errorf("expected signature % X, got % X", SignatureIgnored, v[3])
	}
}
