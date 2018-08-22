package pcsa

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/dgryski/go-metro"
	"github.com/seiflotfy/loglogbeta"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

func RandStringBytesMaskImprSrc(n uint32) string {
	b := make([]byte, n)
	for i := uint32(0); i < n; i++ {
		b[i] = letterBytes[rand.Int()%len(letterBytes)]
	}
	return string(b)
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
	tc, _ := NewTC(16)
	llb := loglogbeta.New()
	step := 1000000
	unique := map[string]bool{}
	for i := 1; len(unique) <= 10000000; i++ {
		str := RandStringBytesMaskImprSrc(rand.Uint32() % 32)
		hash := metro.Hash64Str(str, 1337)
		sk.AddHash(hash)
		tc.AddHash(hash)
		llb.AddHash(hash)
		unique[str] = true

		if len(unique)%step == 0 {
			fmt.Println("---")
			exact := uint64(len(unique))
			res1 := uint64(sk.Cardinality())

			ratio := 100 * estimateError(res1, exact)
			if math.Abs(ratio) > 2 {
				t.Errorf("Normal: Exact %d, got %d which is %.2f%% error", exact, res1, ratio)
			}

			res2 := uint64(tc.Cardinality())
			ratio2 := 100 * estimateError(res2, exact)
			if math.Abs(ratio2) > 2 {
				t.Errorf("TailCut: Exact %d, got %d which is %.2f%% error", exact, res2, ratio2)
			}

			res3 := uint64(llb.Cardinality())
			ratio3 := 100 * estimateError(res3, exact)
			if math.Abs(ratio3) > 2 {
				t.Errorf("HyperLogLog: Exact %d, got %d which is %.2f%% error", exact, res3, ratio3)
			}

			fmt.Printf(">>> %d %2f %2f %2f %d\n", exact, ratio, ratio2, ratio3, tc.bitmaps.base)

		}
	}
}
