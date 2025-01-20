package arrays3

func NextPermutation(A []int) []int {
	n := len(A)
	pivot := -1
	for i := n - 2; i >= 0; i-- {
		if A[i] < A[i+1] {
			pivot = i
			break
		}
	}

	if pivot == -1 {
		reverse(A, 0, n-1)
		return A
	}

	for i := n - 1; i > pivot; i-- {
		if A[i] > A[pivot] {
			A[i], A[pivot] = A[pivot], A[i]
			break
		}
	}

	reverse(A, pivot+1, n-1)
	return A
}

func reverse(A []int, x, y int) {

	for x < y {
		A[x], A[y] = A[y], A[x]
		x++
		y--
	}
}
