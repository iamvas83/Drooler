package intermediateslidingwindow

func SwapNumBers(A []int, B int) int {
	n := len(A)
	if n == 1 {
		return 0
	}
	if B == 0 {
		return 0
	}
	good := 0
	for i := 0; i < n; i++ {
		if A[i] <= B {
			good++
		}
	}

	//window size
	k := good

	bad := 0

	for i := 0; i < k; i++ {
		if A[i] > B {
			bad++
		}
	}

	s := 1
	e := k
	ans := bad

	for e < n {
		out := A[s-1]
		in := A[e]
		if out > B {
			bad--
		}
		if in > B {
			good++
		}
		ans = min(ans, bad)
		s++
		e++
	}
	return ans
}
