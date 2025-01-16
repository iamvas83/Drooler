package arrays1

func RainWaterTrap(A []int) int {
	n := len(A)
	leftMax := make([]int, len(A))
	rightMax := make([]int, len(A))

	leftMax[0] = A[0]
	rightMax[n-1] = A[n-1]

	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], A[i])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], A[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		level := min(leftMax[i], rightMax[i])
		ans += level - A[i]
	}
	return ans
}
