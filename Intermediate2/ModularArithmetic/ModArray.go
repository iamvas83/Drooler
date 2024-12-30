package modulararithmetic

func ModArray(A []int, B int) int {
	ans := 0
	power := 1

	n := len(A)

	for i := n - 1; i > 0; i-- {
		ans = (ans + (A[i]%B*power%B)%B) % B
		power = (power * 10) % B
	}
	return ans
}
