package types

// A Node represents a node from a labeled property graph and consists of a unique identifier
// (within the scope of its origin graph), a list of labels and a map of properties.
// Node (signature=0x4E) {
//     Integer             nodeIdentity
//     List<String>        labels
//     Map<String, Value>  properties
// }
func Node(nodeIdentity Integer, labels List, properties Map) Structure {
	return Structure{
		Fields: []Any{
			nodeIdentity,
			labels,
			properties,
		},
		Signature: SignatureNode,
	}
}
