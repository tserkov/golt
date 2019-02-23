package messages

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/tserkov/golt/types"
)

func TestInitWithCreds(t *testing.T) {
	clientName := types.String("golt/1.0")
	msg := Init(clientName, "user", "password")

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

	// Make sure the message signature is INIT
	if v[3] != SignatureInit {
		t.Errorf("expected signature % X, got % X", SignatureInit, v[3])
	}

	// Make sure first field is a string
	if v[4] != byte(types.MarkerTinyString+8) {
		t.Errorf("expected first field type % X, got % X", byte(types.MarkerTinyString+len(clientName)), v[4])
	}

	// Make sure second field is a 3-field TINY_MAP (scheme, principal, credentials)
	if v[4+len(clientName)+1] != byte(types.MarkerTinyMap+3) {
		t.Errorf("expected second field type % X, got % X", byte(types.MarkerTinyMap+3), v[4+len(clientName)+1])
	}
}

func TestInitNoCreds(t *testing.T) {
	clientName := types.String("golt/1.0")
	msg := Init(clientName, "", "")

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

	// Make sure the message signature is INIT
	if v[3] != SignatureInit {
		t.Errorf("expected signature % X, got % X", SignatureInit, v[3])
	}

	// Make sure first field is a string
	if v[4] != byte(types.MarkerTinyString+8) {
		t.Errorf("expected first field type % X, go % X", byte(types.MarkerTinyString+len(clientName)), v[4])
	}

	// Make sure second field is a 1-field TINY_MAP (scheme)
	if v[4+len(clientName)+1] != byte(types.MarkerTinyMap+1) {
		t.Errorf("expected second field type % X, got % X", byte(types.MarkerTinyMap+1), v[4+len(clientName)+1])
	}
}
