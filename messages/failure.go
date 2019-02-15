package messages

import (
	"github.com/tserkov/golt/types"
)

// The FAILURE message is a server summary message used to signal that a corresponding client message
// has encountered an error while being processed.
//
// See https://boltprotocol.org/v1/#message-failure
func Failure(metadata types.Map) types.Structure {
	return types.Structure{
		Fields: []types.Any{
			metadata,
		},
		Signature: SignatureFailure,
	}
}
