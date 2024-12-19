package intermediate2dmatrices

func RowSum(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])

	ans := make([]int, n)

	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < m; j++ {
			sum += matrix[i][j]
		}
		ans[i] = sum
	}
	return ans
}
