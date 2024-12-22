package intermediatecarryforward

import "math"

func ClosestMinMax(A []int) int {
	min := math.MinInt
	max := math.MaxInt
	n := len(A)
	for i := 0; i < n; i++ {
		if A[i] < min {
			min = A[i]
		}
		if A[i] > max {
			max = A[i]
		}
	}
	if max == min {
		return 1
	}

	max_i := -1
	min_i := -1

	ans := n

	for i := n - 1; i >= 0; i-- {
		if A[i] == min {
			min_i = i
			if max_i != -1 {
				if ans > max_i-min_i+1 {
					ans = max_i - min_i + 1
				}
			}
		} else if A[i] == max {
			max_i = i
			if min_i != -1 {
				if ans > min_i-max_i+1 {
					ans = min_i - max_i + 1
				}
			}
		}
	}
	return ans

}
