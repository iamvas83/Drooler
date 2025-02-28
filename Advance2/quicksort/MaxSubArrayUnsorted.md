## 

``` go
package main

import (
	"fmt"
)

func subUnsort(A []int) []int {
	N := len(A)
	start, end := -1, -1

	// Find the first element which is out of order from the left
	for i := 0; i < N-1; i++ {
		if A[i] > A[i+1] {
			start = i
			break
		}
	}

	// If no such element is found, array is already sorted
	if start == -1 {
		return []int{-1}
	}

	// Find the first element which is out of order from the right
	for i := N - 1; i > 0; i-- {
		if A[i] < A[i-1] {
			end = i
			break
		}
	}

	// Find the min and max in the subarray A[start:end]
	minVal, maxVal := A[start], A[start]
	for i := start; i <= end; i++ {
		if A[i] < minVal {
			minVal = A[i]
		}
		if A[i] > maxVal {
			maxVal = A[i]
		}
	}

	// Expand start to include any number greater than minVal
	for i := 0; i < start; i++ {
		if A[i] > minVal {
			start = i
			break
		}
	}

	// Expand end to include any number smaller than maxVal
	for i := N - 1; i > end; i-- {
		if A[i] < maxVal {
			end = i
			break
		}
	}

	return []int{start, end}
}

func main() {
	A1 := []int{1, 3, 2, 4, 5}
	A2 := []int{1, 2, 3, 4, 5}

	fmt.Println(subUnsort(A1)) // Output: [1, 2]
	fmt.Println(subUnsort(A2)) // Output: [-1]
}

```