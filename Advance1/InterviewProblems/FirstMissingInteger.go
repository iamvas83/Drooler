package arrays3

func FirstMissingPositive(A []int) int {
	i := 0
	res := 0
	if len(A) == 1 && A[i] > 0 {
		return A[i] + 1
	}
	for i < len(A) {
		if A[i] > 0 && A[i] <= len(A) {
			curr_index := A[i] - 1
			if A[curr_index] != A[i] {
				A[curr_index], A[i] = A[i], A[curr_index]
			} else {
				i++
			}
		} else {
			i++
		}
	}
	//fmt.Println(A)
	for i := 0; i < len(A); i++ {
		res = 0
		if A[i] != i+1 {
			res = i + 1
			break
		} else {
			res = len(A) + 1
		}
	}
	return res
}
