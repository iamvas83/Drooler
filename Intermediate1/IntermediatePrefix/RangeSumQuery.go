package intermediatedsa

/**
 * @input A : Integer array
 * @input B : 2D integer array
 *
 * @Output Long array.
 */
func rangeSum(A []int, B [][]int) []int64 {
	pf := make([]int, len(A))
	pf[0] = A[0]
	for i := 1; i < len(pf); i++ {
		pf[i] = pf[i-1] + A[i]
	}

	res := make([]int64, len(B))
	for i, h := range B {
		l := h[0]
		r := h[1]
		if l == 0 {
			res[i] = int64(pf[r])
		} else {
			res[i] = int64(pf[r]) - int64(pf[l-1])
		}
	}
	return res
}
