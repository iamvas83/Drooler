package intermediatedsa

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func EquillibriumIndex(A []int) int {
	pf := make([]int, len(A))
	pf[0] = A[0]
	for i := 1; i < len(A); i++ {
		pf[i] = pf[i-1] + A[i]
	}

	for i := 0; i < len(A); i++ {
		if i == 0 {
			//sum from i-1 to n-1
			sumR := pf[len(A)-1] - pf[i]
			if sumR == 0 {
				return i
			}
		} else if i == len(A)-1 {
			//sum from i to n-2
			sumL := pf[len(A)-2]
			if sumL == 0 {
				return i
			}
		} else {
			sumL := pf[i-1]
			sumR := pf[len(A)-1] - pf[i]
			if sumL == sumR {
				return i
			}
		}
	}
	return -1

}
