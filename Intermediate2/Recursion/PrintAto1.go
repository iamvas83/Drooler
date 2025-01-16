package recursion

import "fmt"

func PrintAto1(A int) {
	print(A)
	fmt.Println()
}

func print(A int) {
	if A == 1 {
		fmt.Print(A, " ")
		return
	}
	fmt.Print(A, " ")
	print(A - 1)
}
