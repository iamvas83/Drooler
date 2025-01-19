package arrays2

func SearchMatrix_Mod(A [][]int, B int) int {

	i := 0
	j := len(A[0]) - 1
	ans := -1
	for i < len(A) && j >= 0 {
		if A[i][j] > B {
			j--
		} else if A[i][j] < B {
			i++
		} else {

			for k := j; k >= 0; k-- {
				if A[i][k] == B {
					ans = ((i+1)*1009 + (k + 1))
				}
			}
			break
		}
	}
	return ans
}
