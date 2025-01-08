package hashing

func CountElements(A []int, B []int) []int {
	m1 := make(map[int]int)

	res := make([]int, 0)
	for i := 0; i < len(B); i++ {
		if _, ok := m1[B[i]]; !ok {
			m1[B[i]] = 1
		} else {
			m1[B[i]] = m1[B[i]] + 1
		}
	}

	for i := 0; i < len(A); i++ {
		if _, ok := m1[A[i]]; ok {
			cnt := m1[A[i]]
			res = append(res, A[i])
			m1[A[i]] = cnt - 1
			if m1[A[i]] == 0 {
				delete(m1, A[i])
			}
		}
	}
	return res
}
