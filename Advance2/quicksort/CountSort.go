package advance2

/**
 * @input A : Integer array
 *
 * @Output Integer array.
 */
//import "fmt"
import "math"

func CountSort(A []int) []int {
	maxi := math.MinInt
	for i := 0; i < len(A); i++ {
		if A[i] > maxi {
			maxi = A[i]
		}
	}
	m := make(map[int]int)
	temp := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		if _, ok := m[A[i]]; ok {
			m[A[i]] = m[A[i]] + 1
		} else {
			m[A[i]] = 1
		}
	}
	// fmt.Print(m)
	cnt := 0
	for i := 0; i <= maxi; i++ {
		if _, ok := m[i]; ok {

			for j := 0; j < m[i]; j++ {
				temp[cnt] = i
				cnt++
			}
		}
	}
	return temp
}
