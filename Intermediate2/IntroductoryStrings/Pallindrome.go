package introductorystrings

import "fmt"

func IsPallindrome(A string) {
	runes := []rune(A)

	l, r := 0, len(runes)-1
	flag := true
	for l <= r {
		if runes[l] != runes[r] {
			flag = false
			break
		}
		l++
		r--
	}
	if !flag {
		fmt.Printf("String %s is not a pallindrome", A)
	} else {
		fmt.Printf("String %s is a pallindrome", A)
	}

}
