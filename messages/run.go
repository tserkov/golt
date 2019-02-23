package messages

import (
	"github.com/tserkov/golt/types"
)

// The RUN message is a client message used to pass a statement for execution
// on the server.
//
// See https://boltprotocol.org/v1/#message-run
func Run(statement types.String, parameters types.Map) types.Structure {
	return types.Structure{
		Fields: []types.Value{
			statement,
			parameters,
		},
		Signature: SignatureRun,
	}
}
