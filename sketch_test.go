package pcsa

import (
	"math/rand"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func estimateError(got, exp uint64) float64 {
	var delta uint64
	if got > exp {
		delta = got - exp
	} else {
		delta = exp - got
	}
	return float64(delta) / float64(exp)
}

func RandStringBytesMaskImprSrc(n uint32) string {
	b := make([]byte, n)
	for i := uint32(0); i < n; i++ {
		b[i] = letterBytes[rand.Int()%len(letterBytes)]
	}
	return string(b)
}

func TestCardinalityZero(t *testing.T) {
	sk := NewDefault()
	if card := sk.Cardinality(); card != 0 {
		t.Error("exepcted cardinality == 0, got", card)
	}
}

func TestCardinalityLinear(t *testing.T) {
	sk := NewDefault()
	step := 100000
	unique := map[string]bool{}

	for i := 1; len(unique) <= 10000000; i++ {
		str := RandStringBytesMaskImprSrc(rand.Uint32() % 32)
		sk.Add([]byte(str))
		unique[str] = true

		if len(unique)%step == 0 {
			exact := uint64(len(unique))
			res := uint64(sk.Cardinality())
			step *= 10

			ratio := 100 * estimateError(res, exact)
			if ratio > 2 {
				t.Errorf("Exact %d, got %d which is %.2f%% error", exact, res, ratio)
			}

		}
	}
}
