package mnemonic

// Mnemonic is an enum type for opcode mnemonics.
type Mnemonic uint8

const (
	// SYS_addr is not implemented
	SYS_addr   = Mnemonic(iota)
	LD_Vx_byte = Mnemonic(iota)
	LD_Vx_Vy
	LD_Vx_DT
	LD_Vx_K
	LD_Vx_I
	LD_DT_Vx
	LD_ST_Vx
	LD_F_Vx
	LD_B_Vx
	LD_I_Vx
	LD_I_addr
)
