package main

import (
	"fmt"
	"math/rand"
	"time"

	pcsa "github.com/seiflotfy/fm85"
	"github.com/seiflotfy/loglogbeta"
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

func main() {
	sk := pcsa.NewDefault()
	llb := loglogbeta.New()
	rand.Seed(time.Now().Unix())
	step := 1000000
	unique := map[uint64]bool{}
	for i := 1; len(unique) < 100000000; i++ {
		hash := rand.Uint64()
		sk.AddHash(hash)
		llb.AddHash(hash)
		unique[hash] = true

		if len(unique)%step == 0 {
			exact := uint64(len(unique))

			res1 := uint64(sk.Cardinality())
			ratio1 := 100 * estimateError(res1, exact)

			res2 := uint64(llb.Cardinality())
			ratio2 := 100 * estimateError(res2, exact)

			fmt.Printf(">>> Exact: %d\tTailCut: %.4f\tLogLogBeta: %.4f\n", exact, ratio1, ratio2)
		}
	}
}
