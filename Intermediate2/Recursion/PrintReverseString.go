package recursion

import "fmt"

func PrintReverseString(A string) {
	i := 0
	j := len(A) - 1
	runes := []rune(A)
	reverse(runes, i, j)
	fmt.Printf("reversed string %s", string(runes))
}

func reverse(x []rune, i, j int) {
	if i > j {
		return
	}

	x[i], x[j] = x[j], x[i]

	reverse(x, i+1, j-1)

}
