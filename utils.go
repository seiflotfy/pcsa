package pcsa

func consecutiveDivide(uniformHashValue uint64, div uint64) uint64 {
	var res uint64
	for uniformHashValue%div != 0 {
		uniformHashValue /= div
		res++
	}
	return res
}

func gdHash(x uint64, div uint64) uint64 {
	return consecutiveDivide(x, div)
}
