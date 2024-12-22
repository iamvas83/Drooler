package intermediate2dmatrices

func RotateMatrix(A [][]int) {
	n := len(A)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			A[i][j], A[j][i] = A[j][i], A[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			A[i][j], A[i][n-j-1] = A[i][n-j-1], A[i][j]
		}
	}
}
