package intermediate2dmatrices

func DiagonalMatrix(A [][]int) int {
	sum := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			if i == j {
				sum += A[i][j]
			}
		}
	}
	return sum
}

// Also
func SumOfDiagonal(A [][]int) int {
	sum := 0
	for i, h := range A {
		for j := range h {

			if i == j {
				sum += A[i][j]
			}

		}
	}
	return sum
}
