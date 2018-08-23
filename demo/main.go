package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/seiflotfy/loglogbeta"
	"github.com/seiflotfy/pcsa"
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
	step := 100000000
	for i := 1; i < 1000000000; i++ {
		hash := rand.Uint64()
		sk.AddHash(hash)
		llb.AddHash(hash)

		if i%step == 0 {
			exact := uint64(i)

			res1 := uint64(sk.Cardinality())
			ratio1 := 100 * estimateError(res1, exact)

			res2 := uint64(llb.Cardinality())
			ratio2 := 100 * estimateError(res2, exact)

			fmt.Printf("Exact Cardinality: %d\tPCSA-TailCut (%%err): %.4f\tLogLogBeta (%%err): %.4f\n", exact, ratio1, ratio2)
		}
	}
}
