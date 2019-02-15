package messages

import (
	"github.com/tserkov/golt/types"
)

// The DISCARD_ALL message is a client message used to discard all remaining items
// from the active result stream.
//
// See https://boltprotocol.org/v1/#message-discard-all
func DiscardAll() types.Structure {
	return types.Structure{
		Signature: SignatureDiscardAll,
	}
}
