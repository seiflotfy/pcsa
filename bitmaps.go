package pcsa

import (
	"math/bits"
)

const maxLZ uint8 = 8

// TailCutBitmap ...
type TailCutBitmap struct {
	registers []uint8
	base      uint8
}

// NewTailCutBitmap ...
func NewTailCutBitmap(k uint64) *TailCutBitmap {
	return &TailCutBitmap{
		registers: make([]uint8, k),
		base:      0,
	}
}

// MinLZ ...
func (tcb *TailCutBitmap) MinLZ() uint8 {
	min := maxLZ
	for _, val := range tcb.registers {
		if lz := uint8(bits.TrailingZeros8(^val)); lz < min {
			min = lz
		}
	}
	return min
}

func (tcb *TailCutBitmap) shift(offset uint8) {
	for i := range tcb.registers {
		tcb.registers[i] >>= offset
	}
}

// Flip ...
func (tcb *TailCutBitmap) Flip(i uint64, lz uint8) {
	if lz < tcb.base {
		return
	}
	if lz-tcb.base >= maxLZ {
		tcb.rebase()
	}

	diff := lz - tcb.base
	if c1 := maxLZ - 1; c1 < diff {
		diff = c1
	}
	newVal := uint8(1 << diff)
	tcb.registers[i] |= newVal
}

// LZ ...
func (tcb *TailCutBitmap) LZ(i uint64) uint8 {
	return uint8(bits.TrailingZeros8(^tcb.registers[i])) + tcb.base
}

// Rebase ...
func (tcb *TailCutBitmap) rebase() {
	if min := tcb.MinLZ(); min > 0 {
		for i := range tcb.registers {
			tcb.registers[i] >>= min
		}
		tcb.base += min
	}
}
