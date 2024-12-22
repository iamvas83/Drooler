package intermediate2dmatrices

func AntiDiagonal(A [][]int) [][]int {
	N := len(A)
	result := make([][]int, 2*N-1)

	for i := range result {
		result[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			result[i+j][min(i, j)] = A[i][j]
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
