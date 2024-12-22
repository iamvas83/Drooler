package intermediatedsainterviewproblemsarrays

func SpecialIndex(A []int) int {
	pfeven := make([]int, len(A))
	pfodd := make([]int, len(A))
	pfeven[0] = A[0]
	pfodd[0] = 0

	for i := 1; i < len(A); i++ {
		if i%2 == 0 {
			pfeven[i] = pfeven[i-1] + A[i]
			pfodd[i] = pfodd[i-1]
		} else {
			pfodd[i] = pfodd[i-1] + A[i]
			pfeven[i] = pfeven[i-1]
		}
	}

	ans := 0
	sumeven := 0
	sumodd := 0
	if pfeven[len(A)-1] == pfodd[len(A)-1] {
		ans++
	}
	for i := 1; i < len(A); i++ {

		sumeven = pfeven[i-1] + pfodd[len(A)-1] - pfodd[i]

		sumodd = pfodd[i-1] + pfeven[len(A)-1] - pfeven[i]

		if sumeven == sumodd {
			ans++
		}
	}

	return ans
}
