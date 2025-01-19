package arrays2

type index struct {
	i int
	j int
}

func Row2Col(A [][]int) [][]int {
	count := 0
	n := len(A)
	m := len(A[0])
	idx := map[int]index{}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == 0 {
				idx[count] = index{i, j}
				count++
			}
		}
	}
	for i := 0; i < len(idx); i++ {
		if _, ok := idx[i]; ok {
			val := idx[i]

			for k := 0; k < n; k++ {
				A[val.i][k] = 0
			}
			for k := 0; k < m; k++ {
				A[k][val.j] = 0
			}

		}
	}
	return A
}
