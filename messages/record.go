package messages

import (
	"github.com/tserkov/golt/types"
)

// The RECORD message is a server detail message used to deliver data from the server to the client.
// Each record message contains a single List, which in turn contains the fields of the record in order.
//
// See https://boltprotocol.org/v1/#message-record
func Record(fields types.List) types.Structure {
	return types.Structure{
		Fields: []types.Any{
			fields,
		},
		Signature: SignatureRecord,
	}
}
