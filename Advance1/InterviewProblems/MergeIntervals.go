package arrays3

import "sort"

type Interval struct {
	start, end int
}

func Merge(A []Interval) []Interval {
	if len(A) < 2 {
		return A
	}
	sort.SliceStable(A, func(i, j int) bool {
		return A[i].start < A[j].start
	})

	first := 0
	for second := 1; second < len(A); second++ {
		if A[second].start <= A[first].end {
			A[first].end = max(A[first].end, A[second].end)
		} else {
			first++
			A[first].start = A[second].start
			A[first].end = A[second].end
		}
	}
	var ans []Interval
	for i := 0; i <= first; i++ {
		ans = append(ans, A[i])
	}
	return ans
}
