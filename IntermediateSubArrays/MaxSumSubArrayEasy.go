package intermediatesubarrays

import "fmt"

func MaxSumSubArray(A []int, sum int, size int) {
	ans := 0
	maxSum := 0
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			ans += A[j]
			if ans <= sum {
				if maxSum < ans {
					maxSum = ans
				}
			} else {
				break
			}
		}
	}
	fmt.Printf("max sum sub array %v\n", maxSum)
}
