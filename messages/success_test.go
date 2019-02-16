package messages

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/tserkov/golt/types"
)

func TestSuccess(t *testing.T) {
	msg := Success(types.Map{})

	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	if err := msg.Serialize(w); err != nil {
		t.Error(err)
		return
	}
	w.Flush()

	v := buf.Bytes()

	// Make sure the message is a 1-field struct
	if v[2] != byte(types.MarkerTinyStruct+1) {
		t.Errorf("expected message type % X, got % X", byte(types.MarkerTinyStruct+1), v[2])
	}

	// Make sure the message signature is SUCCESS
	if v[3] != SignatureSuccess {
		t.Errorf("expected signature % X, got % X", SignatureSuccess, v[3])
	}
}
