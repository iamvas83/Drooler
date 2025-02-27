package advance2

/**
 * @input A : Integer array
 * @input B : Integer array
 *
 * @Output Integer array.
 */
func Merge2SorteedArrays(A []int, B []int) []int {
	m := len(A)
	n := len(B)
	temp := make([]int, m+n)
	cnt := 0
	i, j := 0, 0

	for i < m && j < n {
		if A[i] <= B[j] {
			temp[cnt] = A[i]
			cnt++
			i++
		} else {
			temp[cnt] = B[j]
			cnt++
			j++
		}
	}

	for i < m {
		temp[cnt] = A[i]
		cnt++
		i++
	}
	for j < n {
		temp[cnt] = B[j]
		cnt++
		j++
	}
	return temp
}
