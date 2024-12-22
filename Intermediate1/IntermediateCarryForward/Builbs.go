package intermediatecarryforward

import "fmt"

func NumberofSwitches(A []int) {
	fmt.Println(A)
	cnt := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			cnt++
			A[i] = 1
			for j := i; j < len(A); j++ {
				A[j] = 1 - A[j]
			}
		}
	}
	fmt.Println(cnt)
}

func NumberofSwitches2(A []int) {

	cnt := 0
	for i := 0; i < len(A); i++ {
		state := A[i]
		if state%2 == 0 {
			state = A[i]
		} else {
			state = 1 - A[i]
		}

		if state == 0 {
			cnt++
		}
	}
	fmt.Print(cnt)
}
