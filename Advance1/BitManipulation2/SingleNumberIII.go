package bitmanipulation2

func SingleNumberIII(A []int) []int {
	xor := 0
	for i := 0; i < len(A); i++ {
		xor ^= A[i]
	}

	pos := 0

	for i := 0; i < 31; i++ {
		if pos&(1<<i) > 0 {
			pos = i
			break
		}
	}

	set := 0
	unset := 0
	ans := make([]int, 2)
	for i := 0; i < len(A); i++ {
		if A[i]&(1<<pos) > 0 {
			set ^= A[i]
		} else {
			unset ^= A[i]
		}
	}

	if set < unset {
		ans[0] = set
		ans[1] = unset
	} else {
		ans[0] = unset
		ans[1] = set
	}

	return ans

}
