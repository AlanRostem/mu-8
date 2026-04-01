package mu8

// Byte is an alias for uint8
type Byte uint8

// DByte is short for "double-byte"
type DByte uint16

func (db DByte) Nibble(n uint8) uint8 {
	if n > 3 {
		panic("nibble index too high")
	}
	return uint8((db >> (n * 4)) & 0xF)
}
