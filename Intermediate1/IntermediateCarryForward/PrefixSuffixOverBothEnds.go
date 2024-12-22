package intermediatecarryforward

func MaxSumByB(A []int, size int) int {
	n := len(A)
	pf := make([]int, n)
	sf := make([]int, n)

	pf[0] = A[0]

	for i := 1; i < n; i++ {
		pf[i] = pf[i-1] + A[i]
	}

	sf[n-1] = A[n-1]

	for i := n - 2; i >= 0; i-- {
		sf[i] = sf[i+1] + A[i]
	}
	ans := max(pf[n-1], sf[n-size])

	for i := 0; i < size; i++ {
		sum := max(pf[i-1], sf[n-size+i])
		ans = max(sum, ans)
	}
	return ans
}

func max(A, B int) int {
	if A > B {
		return A
	}
	return B
}
