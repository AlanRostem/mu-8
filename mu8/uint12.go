package mu8

import "fmt"

const Uint12Max = 0x1000

type Uint12 struct {
	value uint16
}

func NewUint12(value uint) Uint12 {
	return Uint12{
		value: uint16(value % Uint12Max),
	}
}

func (u *Uint12) Add(value uint) {
	var temp = uint(u.value) + value
	u.value = uint16(temp % Uint12Max)
}

func (u *Uint12) Value() uint {
	return uint(u.value)
}

func (u *Uint12) Set(value uint) {
	u.value = uint16(value % Uint12Max)
}

func (u *Uint12) String() string {
	return fmt.Sprintf("0x%04X", u.value)
}

var _ fmt.Stringer = (*Uint12)(nil)
