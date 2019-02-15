package messages

import (
	"github.com/tserkov/golt/types"
)

// The RESET message is a client message used to return the current session to a "clean" state.
// It will cause the session to IGNORE any message it is currently processing, as well as any message
// before RESET that had not yet begun processing. This allows RESET to abort long-running operations.
// It also means clients must be careful about pipelining RESET. Only send this if you are not currently
// waiting for a result from a prior message, or if you want to explicitly abort any prior message.
//
// See https://boltprotocol.org/v1/#message-reset
func Reset() types.Structure {
	return types.Structure{
		Signature: SignatureReset,
	}
}
