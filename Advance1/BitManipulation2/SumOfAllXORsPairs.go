package bitmanipulation2

func SumOfAllXORPairs(A []int) int {
	const MOD = 1000000007
	n := len(A)
	result := 0

	for i := 0; i < 32; i++ {
		count1 := 0

		for _, num := range A {
			if num&(1<<i) > 0 {
				count1++
			}
		}

		count0 := n - count1
		contribution := (count1 * count0 % MOD) * (1 << i) % MOD
		result = (result + contribution) % MOD
	}
	return result
}
