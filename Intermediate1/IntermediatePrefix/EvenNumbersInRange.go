package intermediatedsa

/**
 * @input A : Integer array
 * @input B : 2D integer array
 *
 * @Output Integer array.
 */
func EVenNumbersInRange(A []int, B [][]int) []int {
	pf := make([]int, len(A))

	if A[0]%2 == 0 {
		pf[0] = 1
	} else {
		pf[0] = 0
	}
	for i := 1; i < len(A); i++ {
		if A[i]%2 == 0 {
			pf[i] = pf[i-1] + 1
		} else {
			pf[i] = pf[i-1]
		}
	}

	res := make([]int, len(B))
	for i, h := range B {
		l := h[0]
		r := h[1]

		if l == 0 {
			res[i] = pf[r]
		} else {
			res[i] = pf[r] - pf[l-1]
		}
	}

	return res
}
