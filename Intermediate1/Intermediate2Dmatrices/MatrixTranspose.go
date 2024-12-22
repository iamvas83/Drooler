package intermediate2dmatrices

func MatrixTranspose(A [][]int) [][]int {
	row := len(A)
	col := len(A[0])
	ans := make([][]int, col)
	for i := range ans {
		ans[i] = make([]int, row)
	}

	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			ans[i][j] = A[j][i]
		}
	}
	return ans
}
