package recursion

func SolveKthSymbol(A int, B int) int {
	res := Ath_row(A)
	return res[B]
}

func Ath_row(A int) []int {
	if A == 1 {
		v1 := make([]int, 0)
		v1 = append(v1, 0)
		return v1
	}

	v1 := Ath_row(A - 1)
	v2 := make([]int, 0)
	for i := 0; i < len(v1); i++ {
		if v1[i] == 0 {
			v2 = append(v2, 0)
			v2 = append(v2, 1)
		} else {
			v2 = append(v2, 1)
			v2 = append(v2, 0)
		}
	}
	return v2

}
