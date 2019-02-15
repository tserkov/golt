package types

// A Relationship represents a relationship from a labeled property graph and consists of a unique
// identifier (within the scope of its origin graph), identifiers for the start and end nodes of
// that relationship, a type and a map of properties.
// Relationship (signature=0x52) {
//     Integer             relIdentity
//     Integer             startNodeIdentity
//     Integer             endNodeIdentity
//     String              type
//     Map<String, Value>  properties
// }
func Relationship(relIdentity, startNodeIdentity, endNodeIdentity Integer, T String, properties Map) Structure {
	return Structure{
		Fields: []Any{
			relIdentity,
			startNodeIdentity,
			endNodeIdentity,
			T,
			properties,
		},
		Signature: SignatureRelationship,
	}
}
