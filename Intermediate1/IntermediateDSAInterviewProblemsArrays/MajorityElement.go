package intermediatedsainterviewproblemsarrays

func MajorityElement(A []int) int {
	me := A[0]
	count := 1
	for i := 1; i < len(A); i++ {
		if A[i] == me {
			count++
		} else {
			count--
		}
		if count == 0 {
			me = A[i]

		}
	}

	count = 0
	for i := 0; i < len(A); i++ {
		if A[i] == me {
			count++
		}
	}

	n := len(A)
	if count > n/2 {
		return me
	}
	return -1

}
