package advance2

import "sort"

func UniqueElements(A []int) int {
	cnt := 0
	i := 1
	j := 0
	sort.Ints(A)
	for j < len(A) {
		if i >= A[j] {
			cnt += i - A[j]
			i++
			j++
		} else {
			i = A[j] + 1
			j++
		}
	}
	return cnt
}
