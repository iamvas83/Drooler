package bitmanipulation2

func MaxANDPair(A []int) int {
	ans := 0
	for i := 31; i >= 0; i-- {
		cnt := 0

		for j := 0; j < len(A); j++ {
			if A[j]&(1<<i) > 0 {
				cnt++
			}
		}

		if cnt >= 2 {
			ans |= (1 << i)
			for j := 0; j < len(A); j++ {
				if A[j]&(1<<i) < 1 {
					A[i] = 0
				}
			}
		}
	}
	return ans
}
