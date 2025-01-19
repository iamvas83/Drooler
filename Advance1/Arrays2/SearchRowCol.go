package arrays2

func SearchMatrix(A [][]int, target int) bool {
	i := 0
	n := len(A)
	j := len(A[0]) - 1

	for i < n && j >= 0 {
		if target == A[i][j] {
			return true
		} else if A[i][j] < target {
			i++
		} else {
			j--
		}
	}

	return false
}
