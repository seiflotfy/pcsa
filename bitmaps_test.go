package pcsa

import (
	"fmt"
	"testing"
)

func TestTailCutBitmapMax(t *testing.T) {
	tcb := NewTailCutBitmap(4)
	for i := range tcb.registers {
		tcb.registers[i] = uint8((i + 1) * 4)
		fmt.Printf("%08b\n", tcb.registers[i])
	}
	fmt.Println(tcb.registers)
	fmt.Println(tcb.MinLZ())
	tcb.shift(1)
	for i := range tcb.registers {
		fmt.Printf("%08b\n", tcb.registers[i])
	}
}

func TestTailCutBitmapSet(t *testing.T) {
	tcb := NewTailCutBitmap(4)
	tcb.registers = []uint8{129, 1, 3, 3}

	fmt.Println("---")
	for i := range tcb.registers {
		fmt.Printf("%08b (%d)\n", tcb.registers[i], tcb.registers[i])
	}

	tcb.Flip(3, 9)
	fmt.Println("---")
	for i := range tcb.registers {
		fmt.Printf("%08b   %d\n", tcb.registers[i], tcb.LZ(uint64(i)))
	}
}
