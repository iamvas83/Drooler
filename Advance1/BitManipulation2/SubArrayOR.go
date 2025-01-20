package bitmanipulation2

func SubarrayOR(A []int) int {
	n := len(A)
	id := make([]int, 32)
	ans := 0
	for i := 1; i <= n; i++ {
		tmp := A[i-1]
		for j := 0; j <= 31; j++ {
			pw := (1 << j)
			if (tmp & pw) != 0 {
				ans += pw * i
				id[j] = i
			} else if id[j] != 0 {
				ans += pw * id[j]
			}
		}
	}
	return ans % 1000000007
}
