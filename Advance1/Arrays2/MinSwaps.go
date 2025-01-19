package arrays2

func MinSwaps(A []int, B int) int {
	good := 0

	for i := 0; i <= len(A); i++ {
		if A[i] <= B {
			good++
		}
	}

	k := good
	bad := 0
	for i := 0; i < k; i++ {
		bad++
	}
	e := k
	s := 1
	ans := bad

	for i := 1; i < len(A); i++ {
		if A[e] > B {
			bad++
		}
		if A[s-1] > B {
			bad--
		}
		ans = min(ans, bad)
		e++
		s--
	}
	return ans
}
