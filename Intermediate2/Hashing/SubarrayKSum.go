package hashing

func SubArrayK(A []int, sum int) int {
	m := map[int]int{}
	m[0] = 1
	psum := 0

	res := 0
	for i := 0; i < len(A); i++ {
		psum += A[i]

		if _, ok := m[psum-sum]; ok {
			res += m[psum-sum]
		}
		m[psum] = m[psum] + 1
	}
	return res
}
