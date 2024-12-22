package intermediatedsainterviewproblemsarrays

func SumofOddIndices(A []int, B [][]int) []int {
	pf := make([]int, len(A))
	pf[0] = 0
	ans := make([]int, len(B))
	for i := 1; i < len(A); i++ {
		if i%2 != 0 {
			pf[i] = pf[i-1] + A[i]
		} else {
			pf[i] = pf[i-1]
		}

	}
	//fmt.Println(pf)
	//count:=0
	for index, a := range B {
		m := a[0]
		n := a[1]
		// fmt.Println(pf[m-1],pf[n])
		if m == 0 {
			ans[index] = pf[n]
		} else {
			ans[index] = pf[n] - pf[m-1]
		}

		//count++
	}
	return ans
}
