package messages

import (
	"github.com/tserkov/golt/types"
)

// The ACK_FAILURE message is a client message used to acknowledge a failure the server has sent.
//
// See https://boltprotocol.org/v1/#message-ack-failure
func AckFailure() types.Structure {
	return types.Structure{
		Signature: SignatureAckFailure,
	}
}
