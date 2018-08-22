package pcsa

import (
	"fmt"
	"math"
	"math/bits"

	"github.com/dgryski/go-metro"
)

const (
	phi   = 0.77351
	kappa = 1.75
)

// Sketch ...
type Sketch struct {
	b       uint8
	m       uint64
	bitmaps []uint64
}

// New ...
func New(b uint8) (*Sketch, error) {
	if b < 4 || b > 16 {
		return nil, fmt.Errorf("expected 4 <= b <= 16, got %v", b)
	}
	m := uint64(1) << b
	return &Sketch{
		b:       b,
		m:       m,
		bitmaps: make([]uint64, m),
	}, nil
}

// NewDefault ...
func NewDefault() *Sketch {
	sk, _ := New(14)
	return sk
}

// Add ...
func (sk *Sketch) Add(val []byte) {
	sk.AddHash(metro.Hash64(val, 1337))
}

// AddHash ...
func (sk *Sketch) AddHash(x uint64) {
	idx := x >> (64 - sk.b)
	lz := bits.TrailingZeros64(x)
	sk.bitmaps[idx] |= 1 << uint64(lz)
}

// Cardinality ...
func (sk *Sketch) Cardinality() uint64 {
	sum := float64(0)

	for _, val := range sk.bitmaps {
		sum += float64(bits.TrailingZeros64(^val))
	}

	m := float64(sk.m)
	res := m / phi * (math.Pow(2, float64(sum)/m) - math.Pow(2, -kappa*sum/m))
	return uint64(res + 0.5)
}
