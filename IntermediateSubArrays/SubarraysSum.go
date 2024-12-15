package intermediatesubarrays

import "fmt"

func SubarraysSum2(A []int) {
	sum := 0
	for i := 0; i < len(A); i++ {
		for j := i; j < len(A); j++ {

			for k := i; k <= j; k++ {
				sum += A[i]
			}
		}
	}

	fmt.Printf("Subarrays sum %v\n", sum)
}

func SubarraysSum(A []int) {
	n := len(A)
	sum := 0
	for i := 0; i < n; i++ {
		sum += A[i] * (i + 1) * (n - i)
	}
	fmt.Printf("Subarrays sum %v\n", sum)
}
