package bitmanipulation2

func Find2MissingNumbers(A []int) []int {
	xor := 0
	n := len(A)
	for i := 0; i < n; i++ {
		xor ^= A[i]
	}

	for i := 1; i < n+2; i++ {
		xor ^= i
	}

	pos := 0
	for i := 0; i < 31; i++ {
		if xor&(1<<i) > 0 {
			pos = i
			break
		}
	}
	ans := make([]int, 2)

	set := 0
	unset := 0

	for i := 0; i < n; i++ {
		if A[i]&(1<<pos) > 0 {
			set ^= A[i]
		} else {
			unset ^= A[i]
		}
	}
	for i := 1; i < n+2; i++ {
		if i&(1<<pos) > 0 {
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
