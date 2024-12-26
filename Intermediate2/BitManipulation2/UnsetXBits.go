package bitmanipulation2

func UnsetXBits(A, B int) int {
	ans := A
	for i := 0; i < B; i++ {

		if (A & (1 << i)) != 0 {
			ans -= 1 << i
		}
	}

	return ans
}
