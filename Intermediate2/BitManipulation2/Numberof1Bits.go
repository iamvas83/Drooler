package bitmanipulation2

func CountBits(A int) int {
	cnt := 0
	for A > 0 {
		n := A % 2
		if n == 1 {
			cnt++
		}
		A /= 2
	}
	return cnt
}

func NumSetBits(A int) int {
	cnt := 0

	for A > 0 {
		if A&1 != 0 {
			cnt++
		}
		A = A >> 1
	}
	return cnt
}
