package bitmanipulation

func SingleNumber(A []int) int {
	a := A[0]
	for i := 1; i < len(A); i++ {
		a ^= A[i]
	}
	return a
}
