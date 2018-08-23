package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/seiflotfy/loglogbeta"
	"github.com/seiflotfy/pcsa"
)

func main() {
	sk := pcsa.NewDefault()
	llb := loglogbeta.New()
	rand.Seed(time.Now().Unix())
	step := 5000000
	wins := 0
	loss := 0
	for i := 1; i < 100000000; i++ {
		hash := rand.Uint64()
		sk.AddHash(hash)
		llb.AddHash(hash)

		if i%step == 0 {
			exact := float64(i)

			res1 := float64(sk.Cardinality())
			ratio1 := exact / res1
			delta1 := math.Abs(exact - res1)

			res2 := float64(llb.Cardinality())
			ratio2 := exact / res2
			delta2 := math.Abs(exact - res2)

			if delta1 < delta2 {
				wins++
			} else {
				loss++
			}

			fmt.Printf("Exact Cardinality: %0.f\tPCSA-TailCut (ratio): %.4f\tLogLogBeta (ratio): %.4f\tPCSA-TailCut win: %v\n", exact, ratio1, ratio2, delta1 < delta2)
		}
	}

	fmt.Println("wins:", wins, "loss:", loss)
}
