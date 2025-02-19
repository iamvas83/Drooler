// You can edit this code!
// Click here and start typing.
package arrays3

import "fmt"

func Main_iju() {

	A := []int{8, 3, 6, 3, 11, 9}

	B := [][]int{{0, 3},
		{3, 5},
		{1, 3},
		{2, 4}}

	fmt.Print(Solve(A, B))
}

func Solve(A []int, B [][]int) []int {
	prefix := make([]int, len(A)+1)
	for i := 0; i < len(A); i++ {
		prefix[i+1] = prefix[i]
		if isComposite(A[i]) {
			prefix[i+1]++
		}
	}
	fmt.Print(prefix)
	ans := make([]int, len(B))
	for i := 0; i < len(B); i++ {
		x := B[i][0]
		y := B[i][1]

		res := prefix[y+1] - prefix[x]
		ans[i] = res

	}
	return ans
}

func isComposite(A int) bool {
	if A < 4 {
		return false
	}
	//count := 0
	for i := 2; i*i <= A; i++ {
		if A%i == 0 {
			return true
		}
	}

	//if count > 2 {
	//	return true
	//} else {
	//	return false
	//}

	return false
}
