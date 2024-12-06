package intermediatecarryforward

func LeadersInArray(nums []int) []int {
	n := len(nums)
	ans := []int{}
	max := nums[n-1]
	ans = append(ans, max)
	for i := n - 2; i >= 0; i-- {
		if nums[i] > max {
			max = nums[i]
			ans = append(ans, max)
		}
	}

	return ans
}
