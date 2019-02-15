package messages

import (
	"github.com/tserkov/golt/types"
)

// The PULL_ALL message is a client message used to retrieve all remaining
// items from the active result stream.
//
// See https://boltprotocol.org/v1/#message-pull-all
func PullAll() types.Structure {
	return types.Structure{
		Signature: SignaturePullAll,
	}
}
