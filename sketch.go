package pcsa

import (
	"math"
	"math/bits"

	"github.com/dgryski/go-metro"
)

const (
	phi   = 0.77350
	kappa = 1.4
)

// Sketch ...
type Sketch struct {
	b       uint8
	m       uint64
	bitmaps *TailCutBitmap
}

// New ...
func New(b uint8) (*Sketch, error) {
	m := uint64(1) << b
	return &Sketch{
		b:       b,
		m:       m,
		bitmaps: NewTailCutBitmap(m),
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
	sk.bitmaps.Flip(idx, uint8(lz))
}

func (sk *Sketch) sum() uint64 {
	sum := uint64(0)

	for i := uint64(0); i < sk.m; i++ {
		sum += uint64(sk.bitmaps.LZ(i))
	}

	// TODO: We are always over estimating, so I am trying to subtract something based on our current base
	//sum -= 1 << sk.bitmaps.base
	return sum

}

// Cardinality ...
func (sk *Sketch) Cardinality() uint64 {
	sum := float64(sk.sum())
	m := float64(sk.m)

	// TODO: Trying another correction here
	res := m/phi*(math.Pow(2, float64(sum)/m)) - math.Pow(2, -kappa*sum/m)
	return uint64(res + 0.5)
}
