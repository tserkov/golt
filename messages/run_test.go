package messages

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/tserkov/golt/types"
)

func TestRun(t *testing.T) {
	msg := Run(types.String(""), types.Map{})

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	if err := msg.Serialize(w); err != nil {
		t.Error(err)
		return
	}
	w.Flush()

	v := buf.Bytes()

	// Make sure the message is a 2-field struct
	if v[2] != byte(types.MarkerTinyStruct+2) {
		t.Errorf("expected message type % X, got % X", byte(types.MarkerTinyStruct+2), v[2])
	}

	// Make sure the message signature is RUN
	if v[3] != SignatureRun {
		t.Errorf("expected signature % X, got % X", SignatureRun, v[3])
	}
}
