package modulararithmatic

func ModuloPairSum(A []int, B int) int {
	cnt := make([]int, B)
	var mod int = 1e9 + 7
	ans := 0
	for i := 0; i < len(A); i++ {
		cnt[A[i]%B]++
	}

	ans += (cnt[0] * (cnt[0] - 1)) / 2
	ans %= mod
	i := 1
	j := B - 1
	for i <= j {
		if i == j {
			ans += (cnt[i] * (cnt[j] - 1)) / 2
			ans %= mod
		} else {
			ans += cnt[i] * cnt[j]
			ans %= mod
		}
		i++
		j--
	}
	return ans
}
