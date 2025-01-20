package bitmanipulation2

import (
	"math"
	"sort"
)

func FindMinXor(A []int) int {
	sort.Ints(A)

	minXor := math.MaxInt32
	for i := 0; i < len(A)-1; i++ {
		xorValue := A[i] ^ A[i+1]
		if xorValue < minXor {
			minXor = xorValue
		}
	}
	return minXor
}
