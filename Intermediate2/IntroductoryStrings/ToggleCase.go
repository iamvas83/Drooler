package introductorystrings

import (
	"fmt"
)

func ToggleCase(A string) {
	for i := 0; i < len(A); i++ {
		if A[i] >= 65 && A[i] <= 90 {
			ch := A[i] + 32
			fmt.Printf("Alphbet in Upper case to lower case for A: %v and %v\n", A, string(ch))
		}

		if A[i] >= 97 && A[i] <= 122 {
			ch := A[i] - 32
			fmt.Printf("Alphabet in lower case to Upper case for A: %v and %v\n", A, string(ch))
		}
	}
}
