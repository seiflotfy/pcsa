package pcsa

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func estimateError(got, exp uint64) float64 {
	var delta uint64
	sign := 1.0
	if got > exp {
		delta = got - exp
	} else {
		delta = exp - got
		sign = -1
	}
	return sign * float64(delta) / float64(exp)
}

func TestCardinalityZero(t *testing.T) {
	sk, _ := New(4)
	if card := sk.Cardinality(); card != 0 {
		t.Error("exepcted cardinality == 0, got", card)
	}
}

func TestCardinality10(t *testing.T) {
	sk, _ := New(14)
	for i := uint64(0); i < 10; i++ {
		sk.AddHash(i)
	}
	if card := sk.Cardinality(); card != 10 {
		t.Error("exepcted cardinality == 10, got", card)
	}
}

/*
func TestCardinalityOne(t *testing.T) {
	sk := NewDefault()
	sk.AddHash(1)
	if card := sk.Cardinality(); card != 1 {
		t.Error("exepcted cardinality == 1, got", card)
	}
}
*/

func TestCardinalityLinear(t *testing.T) {
	sk := NewDefault()
	rand.Seed(time.Now().Unix())
	step := 1000000
	unique := map[uint64]bool{}
	for i := 1; len(unique) < 100000000; i++ {
		hash := rand.Uint64()
		sk.AddHash(hash)
		unique[hash] = true

		if len(unique)%step == 0 {
			exact := uint64(len(unique))
			res1 := uint64(sk.Cardinality())

			ratio1 := 100 * estimateError(res1, exact)
			if math.Abs(ratio1) > 2 {
				t.Errorf("Normal: Exact %d, got %d which is %.2f%% error", exact, res1, ratio1)
			}
		}
	}
}
