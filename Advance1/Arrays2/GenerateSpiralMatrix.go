package arrays2

func SearchGenerateMatrix(A int) [][]int {
	i := 0
	j := 0

	ans := make([][]int, A)

	for i := range ans {
		ans[i] = make([]int, A)
	}

	count := 1

	if A == 1 {
		ans[0][0] = count
		return ans
	}

	for A > 1 {
		for k := 1; k < A; k++ {
			ans[i][j] = count
			count++
			j++
		}
		for k := 0; k < A; k++ {
			ans[i][j] = count
			count++
			i++
		}
		for k := 1; k < A; k++ {
			ans[i][j] = count
			count++
			j--
		}
		for k := 1; k < A; k++ {
			ans[i][j] = count
			count++
			i--
		}
		A = A - 2
		i++
		j++
	}
	return ans
}
