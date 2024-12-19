package intermediate2dmatrices

func MinorDiagonal(A [][]int) int {
	n := len(A)
	r := 0
	c := len(A)
	sum := 0
	for r < n {
		sum += A[r][c]
		r++
		c--
	}
	return sum
}
