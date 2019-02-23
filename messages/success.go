package messages

import (
	"github.com/tserkov/golt/types"
)

// The SUCCESS message is a server summary message used to signal that a corresponding client message
// has been received and actioned as intended. The message contains a map of metadata, the contents
// of which depend on the original request.
//
// See https://boltprotocol.org/v1/#message-success
func Success(metadata types.Map) types.Structure {
	return types.Structure{
		Fields: []types.Value{
			metadata,
		},
		Signature: SignatureSuccess,
	}
}
