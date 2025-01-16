package recursion

import "fmt"

func Print1toA(A int) {
	if A == 1 {
		fmt.Print(A, " ")
		return
	}

	Print1toA(A - 1)
	fmt.Print(A, " ")
}

func Solve(A int) {
	Print1toA(A)
	fmt.Print("\n")
}
