package intermediatedsa

/**
 * @input A : Integer array
 *
 * @Output Integer array.
 */
func InPlacePrefixSum(A []int) []int {
	for i := 1; i < len(A); i++ {
		A[i] = A[i-1] + A[i]
	}
	return A
}
