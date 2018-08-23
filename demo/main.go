package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/seiflotfy/loglogbeta"
	"github.com/seiflotfy/pcsa"
)

func main() {
	sk := pcsa.NewDefault()
	llb := loglogbeta.New()
	rand.Seed(time.Now().Unix())
	step := 10000000
	for i := 1; i < 100000000; i++ {
		hash := rand.Uint64()
		sk.AddHash(hash)
		llb.AddHash(hash)

		if i%step == 0 {
			exact := float64(i)

			res1 := float64(sk.Cardinality())
			ratio1 := exact / res1

			res2 := float64(llb.Cardinality())
			ratio2 := exact / res2

			fmt.Printf("Exact Cardinality: %0.f\tPCSA-TailCut (ratio): %.4f\tLogLogBeta (ratio): %.4f\n", exact, ratio1, ratio2)
		}
	}
}
