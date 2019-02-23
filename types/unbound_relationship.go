package types

// An UnboundRelationship represents a relationship relative to a separately known start point and end point.
// UnboundRelationship (signature=0x72) {
//     Integer             relIdentity
//     String              type          // e.g. "KNOWS"
//     Map<String, Value>  properties    // e.g. {since:1999}
// }
func UnboundRelationship(relIdentity Integer, T String, properties Map) Structure {
	return Structure{
		Fields: []Value{
			relIdentity,
			T,
			properties,
		},
		Signature: SignatureUnboundRelationship,
	}
}
