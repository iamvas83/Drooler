package intermediatecarryforward

import "math"

func MaxProfit(nums []int) int {
	n := len(nums)

	min := math.MaxInt
	profit := 0
	maxProfit := 0
	for i := 0; i < n; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > min {
			profit = nums[i] - min
		}

		if maxProfit < profit {
			maxProfit = profit
		}
	}
	return maxProfit
}
