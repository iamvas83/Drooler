package arrays2

import (
	"fmt"
	"math"
)

func MaxSubMatrixSum(A [][]int) int64 {
	n := len(A)
	m := len(A[0])

	// Initialize the maximum sum to a very small number
	maxSum := math.MinInt64

	// Iterate over all pairs of rows
	for rowStart := 0; rowStart < n; rowStart++ {
		// Create an array to store column-wise sums
		colSum := make([]int, m)

		for rowEnd := rowStart; rowEnd < n; rowEnd++ {
			// Update colSum to include values from the current row
			for col := 0; col < m; col++ {
				colSum[col] += A[rowEnd][col]
			}

			// Apply Kadane's algorithm on the updated colSum array
			maxSum = max(maxSum, kadane(colSum))
		}
	}

	return int64(maxSum)
}

// Kadane's Algorithm to find the maximum sum subarray
func kadane(arr []int) int {
	maxEndingHere := 0
	maxSoFar := math.MinInt64

	for _, val := range arr {
		maxEndingHere += val
		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
		}
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}

	return maxSoFar
}

// Utility function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Main_submatrix() {
	A := [][]int{
		{-83, -73, -70, -61},
		{-56, -48, -13, 4},
		{38, 48, 71, 71},
	}
	fmt.Println(MaxSubMatrixSum(A)) // Output: 228
}
