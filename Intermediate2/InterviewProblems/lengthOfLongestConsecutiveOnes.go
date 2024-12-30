package interviewproblems

func LongestConsec1s(A []int) int {
	n := len(A)
	ans := 0
	for i := 0; i < n; i++ {
		if A[i] == 0 {
			j := i - 1
			cnt := 0
			for j >= 0 && A[j] == 1 {
				cnt++
				j--
			}
			L := cnt

			j = i + 1
			cnt = 0
			for j < n && A[j] == 1 {
				cnt++
				j++
			}
			R := cnt

			ans = max(ans, L+R+1)
		}
	}
	return ans
}

func LongestConsec1s_variation(A []int) int {
	n := len(A)
	ans := 0
	ones := 0
	for i := 0; i < n; i++ {
		if A[i] == 1 {
			ones++
		}
	}
	if ones == n {
		ans = n
	}
	for i := 0; i < n; i++ {
		if A[i] == 0 {
			j := i - 1
			cnt := 0

			for j >= 0 && A[j] == 1 {
				cnt++
				j--
			}
			L := cnt

			cnt = 0

			k := i + 1
			for k < n && A[k] == 1 {
				cnt++
				k++
			}
			R := cnt
			total := 0
			if L+R+1 <= ones {
				total = L + R + 1
			} else {
				total = L + R
			}
			ans = max(ans, total)
		}
	}
	return ans
}
