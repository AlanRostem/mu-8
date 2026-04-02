package decode

type Class uint8

const (
	Class0 = Class(iota)
	Class1
	Class2
	Class3
	Class4
	Class5
	Class6 // Contains LD
	Class7
	Class8 // Contains LD
	Class9
	ClassA // Contains LD
	ClassB
	ClassC
	ClassD
	ClassE
	ClassF // Contains LD
)

type IdentitySize uint8

const (
	// IdentitySizeNone represents instructions with no identity (classes with one instruciton)
	IdentitySizeNone = IdentitySize(iota)
	// IdentitySize1 represents instructions where the identity number is one nibble
	IdentitySize1
	// IdentitySize1 represents instructions where the identity number is two nibbles
	IdentitySize2
	// IdentitySize1 represents instructions where the identity number is three nibbles
	IdentitySize3
)

type ClassInfo struct {
	Args []uint8
	// IdentitySize represents the last n nibbles in the opcode.
	IdentitySize IdentitySize
}

// Class0 is omitted since it is a special case
var Classes = map[Class]ClassInfo{
	Class1: {Args: []uint8{3}, IdentitySize: IdentitySizeNone},
	Class2: {Args: []uint8{3}, IdentitySize: IdentitySizeNone},
	Class3: {Args: []uint8{1, 2}, IdentitySize: IdentitySizeNone},
	Class4: {Args: []uint8{1, 2}, IdentitySize: IdentitySizeNone},
	Class5: {Args: []uint8{1, 1}, IdentitySize: IdentitySize1},
	Class6: {Args: []uint8{1, 2}, IdentitySize: IdentitySizeNone},
	Class7: {Args: []uint8{1, 2}, IdentitySize: IdentitySizeNone},
	Class8: {Args: []uint8{1, 1}, IdentitySize: IdentitySize1},
	Class9: {Args: []uint8{1, 1}, IdentitySize: IdentitySize1},
	ClassA: {Args: []uint8{3}, IdentitySize: IdentitySizeNone},
	ClassB: {Args: []uint8{3}, IdentitySize: IdentitySizeNone},
	ClassC: {Args: []uint8{1, 2}, IdentitySize: IdentitySizeNone},
	ClassD: {Args: []uint8{1, 1, 1}, IdentitySize: IdentitySizeNone},
	ClassE: {Args: []uint8{1}, IdentitySize: IdentitySize2},
	ClassF: {Args: []uint8{1}, IdentitySize: IdentitySize2},
}
