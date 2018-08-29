package pcsa

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestConsecutiveDivide(t *testing.T) {
	x := map[uint64]uint64{}
	for i := 0; i < 100000; i++ {
		x[gdHash(rand.Uint64(), 2)]++
	}
	sorted := make([][2]uint64, 0, len(x))
	for k, v := range x {
		sorted = append(sorted, [2]uint64{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
	}) //sort by key
	fmt.Println(sorted)
}
