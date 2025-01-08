package hashing

func DistinctKNums(A []int, B int) []int {
	m := map[int]int{}
	res := make([]int, len(A)-B+1)
	count := 0
	for i := 0; i < B; i++ {
		if _, ok := m[A[i]]; ok {
			m[A[i]] = m[A[i]] + 1
		} else {
			m[A[i]] = 1
		}
	}

	res[count] = len(m)
	count++
	//slide a window
	j := B
	n := len(A)
	i := 1

	for j < n {
		incoming := A[j]
		outgoing := A[i-1]

		if _, ok := m[incoming]; ok {
			m[incoming] = m[incoming] + 1
		} else {
			m[incoming] = 1
		}
		m[outgoing] = m[outgoing] - 1

		if m[outgoing] == 0 {
			delete(m, outgoing)
		}
		res[count] = len(m)
		count++
		i++
		j++
	}
	return res
}
