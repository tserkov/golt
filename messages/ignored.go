package messages

import (
	"github.com/tserkov/golt/types"
)

// The IGNORED message is a server summary message used to signal that a corresponding client message
// has been ignored and not actioned.
//
// See https://boltprotocol.org/v1/#message-ignored
func Ignored() types.Structure {
	return types.Structure{
		Signature: SignatureIgnored,
	}
}
