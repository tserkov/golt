package types

// https://boltprotocol.org/v1/#marker_table
// Every value begins with a marker byte. The marker contains information on
// data type as well as direct or indirect size information for those types that
// require it. How that size information is encoded varies by marker type.
//
// Some values, such as true, can be encoded within a single marker byte and many
// small integers (specifically between -16 and +127) are also encoded within a
// single byte.
//
// A number of marker bytes are reserved for future expansion of the format itself.
// These bytes should not be used, and encountering them in a stream should treated
// as an error.
const (
	MarkerTinyInt    = 0x00 // 0x00..0x7F +TINY_INT
	MarkerTinyString = 0x80 // 0x80..0x8F TINY_STRING
	MarkerTinyList   = 0x90 // 0x90..0x9F TINY_LIST
	MarkerTinyMap    = 0xA0 // 0xA0..0xAF TINY_MAP
	MarkerTinyStruct = 0xB0 // 0xB0..0xBF TINY_STRUCT

	MarkerNull    = 0xC0 // NULL
	MarkerFloat64 = 0xC1 // FLOAT_64
	MarkerFalse   = 0xC2 // FALSE
	MarkerTrue    = 0xC3 // TRUE

	// 0xC4..0xC7 Reserved

	MarkerInt8  = 0xCB // INT_8
	MarkerInt16 = 0xC9
	MarkerInt32 = 0xCA
	MarkerInt64 = 0xCB

	// 0xCC..0xCF Reserved

	MarkerString8  = 0xD0 // STRING_8
	MarkerString16 = 0xD1 // STRING_16
	MarkerString32 = 0xD2 // STRING_32

	// 0xD3 Reserved

	MarkerList8  = 0xD4 // LIST_8
	MarkerList16 = 0xD5 // LIST_16
	MarkerList32 = 0xD6 // LIST_32

	// 0xD7 Reserved

	MarkerMap8  = 0xD8 // MAP_8
	MarkerMap16 = 0xD9 // MAP_16
	MarkerMap32 = 0xDA // MAP_32

	// 0xDB Reserved

	MarkerStruct8  = 0xDC // STRUCT_8
	MarkerStruct16 = 0xDD // STRUCT_16

	// 0xDE..0xEF Reserved

	// 0xF0..0xFF -TINY_INT
)

var (
	MarkerEnd = []byte{0x00, 0x00}
)
