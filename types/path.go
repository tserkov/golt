package types

// A Path is a sequence of alternating nodes and relationships corresponding to a walk in the graph.
// The path always begins and ends with a node. Its representation consists of a list of distinct nodes,
// a list of distinct relationships and a sequence of integers describing the path traversal.
// Path (signature=0x50) {
//     List<Node> nodes
//     List<UnboundRelationship> relationships
//     List<Integer> sequence
// }
func Path(nodes, relationships, sequence List) Structure {
	return Structure{
		Fields: []Value{
			nodes,
			relationships,
			sequence,
		},
		Signature: SignaturePath,
	}
}
