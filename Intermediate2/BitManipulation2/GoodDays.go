package bitmanipulation2

func GoodDays(A int) int {
	cnt := 0
	for i := 0; i < 32; i++ {
		if (A & (1 << i)) != 0 {
			cnt++
		}
	}
	return cnt
}
