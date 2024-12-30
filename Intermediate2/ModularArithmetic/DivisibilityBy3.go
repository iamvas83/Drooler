package modulararithmetic

func DivisibilityBy3(A []int) bool {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}

	if sum%3 == 0 {
		return true
	} else {
		return false
	}
}
