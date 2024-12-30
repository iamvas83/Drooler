package interviewproblems

func RotateArrays(A []int, B []int) [][]int {
	n := len(A)
	res := make([][]int, n)

	for i := range res {
		res[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		for i := 0; i < len(B); i++ {
			res[j][i] = A[j+i] % n
		}
	}
	return res
}
