package recursion

func Is_Num_Magic(A int) int {
	if A < 10 {
		if A == 1 {
			return 1
		} else {
			return 0
		}
	}

	sum := 0

	for A > 0 {
		sum += A % 10
		A /= 10
	}

	return Is_Num_Magic(sum)
}

func IsMagic(A int) bool {

	val := Is_Num_Magic(A)

	if val == 1 {
		return true
	} else {
		return false
	}
}
