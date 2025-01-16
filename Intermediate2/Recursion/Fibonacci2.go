package recursion

func FindAthFibonacci(A int) int {
	if A <= 1 {
		return A
	}

	return FindAthFibonacci(A-1) + FindAthFibonacci(A-2)
}
