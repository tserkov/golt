package messages

import (
	"github.com/tserkov/golt/types"
)

// The INIT message is a client message used once to initialize the session.
// This message is always the first message the client sends after negotiating protocol version.
// Sending any message other than INIT as the first message to the server will result in a FAILURE.
// As described in Failure handling the client must acknowledge failures using ACK_FAILURE, after which INIT may be reattempted.
//
// See https://boltprotocol.org/v1/#message-init
func Init(clientName, username, password types.String) types.Structure {
	fields := make([]types.Any, 2)

	fields[0] = clientName

	if username != "" {
		fields[1] = types.Map{
			"scheme":      types.String("basic"),
			"principal":   username,
			"credentials": password,
		}
	} else {
		fields[1] = types.Map{
			"scheme": types.String("none"),
		}
	}

	return types.Structure{
		Fields:    fields,
		Signature: SignatureInit,
	}
}
