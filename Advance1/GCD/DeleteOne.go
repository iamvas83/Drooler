package gcd

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func DeleteOne(A []int) int {
	n := len(A)
	pf := make([]int, n)
	sf := make([]int, n)

	pf[0] = A[0]
	for i := 1; i < n; i++ {
		pf[i] = gcd(A[i], pf[i-1])
	}
	sf[n-1] = A[n-1]
	for i := n - 2; i >= 0; i-- {
		sf[i] = gcd(sf[i+1], A[i])
	}

	ans := 0

	for i := 0; i < n; i++ {
		if i == 0 {
			ans = max(ans, sf[i+1])
		} else if i == n-1 {
			ans = max(ans, pf[i-1])
		} else {
			ans = max(ans, gcd(pf[i-1], sf[i+1]))
		}
	}
	return ans
}

func gcd(A int, B int) int {

	if A == 0 {
		return B
	}

	return gcd(B%A, A)
}

func max(A, B int) int {
	if A > B {
		return A
	} else {
		return B
	}
}
