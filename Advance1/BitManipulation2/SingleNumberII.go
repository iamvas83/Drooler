package bitmanipulation2

func SingleNumber(A []int) int {
	ans := 0

	for i := 0; i < 31; i++ {
		cnt := 0
		for j := 0; j < len(A); j++ {
			if A[i]&(1<<i) > 0 {
				cnt++
			}
		}

		if cnt%3 == 1 {
			ans |= 1 << i
		}
	}
	return ans
}
