package intermediateslidingwindow

func SubArrayGivenSumLength(A []int, B int, sum int) int {
	start := 0
	end := 0
	n := len(A)
	curr := 0
	for i := end; i < n; i++ {
		curr += A[i]
		if sum < curr {
			curr -= A[start]
			start++
		}

		if curr == sum {
			if B == i-start+1 {
				return 1
			} else {
				return 0
			}
		}
	}
	return 0

}
