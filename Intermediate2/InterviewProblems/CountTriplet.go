package interviewproblems

func CountTriplet(A []int) int {
	ans := 0
	n := len(A)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if A[i] < A[j] && A[j] < A[k] {
					ans++
				}
			}
		}
	}
	return ans
}

func CountTriplet_Eff(A []int) int {
	n := len(A)
	ans := 0
	for i := 0; i < n; i++ {
		cnt_left := 0
		j := i - 1
		for j >= 0 {
			if A[i] > A[j] {
				cnt_left++
			}
			j--
		}

		cnt_right := 0
		k := i + 1
		for k < n {
			if A[i] < A[k] {
				cnt_right++
			}
			k++
		}

		ans += cnt_left * cnt_right
	}
	return ans
}
