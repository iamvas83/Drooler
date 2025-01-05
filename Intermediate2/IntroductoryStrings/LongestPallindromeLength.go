package introductorystrings

import "fmt"

//Interview question
func LongestPallindromeLength2(A string) {

	ans := 0

	for i := 0; i < len(A); i++ {
		for j := i; j < len(A); j++ {
			l := i
			r := j
			count := 0
			for l <= r {
				if A[l] != A[r] {
					break
				} else {
					count = r - l + 1
				}
				l++
				r--
			}
			ans = max(ans, count)
		}
	}

	fmt.Printf("Longest Pallindrome Substring in %s is %d", A, ans)
}

func LongestPallindromeLength1(A string) int {
	ans := 0
	for i := 0; i < len(A); i++ {
		ans = max(ans, expand(A, i, i))
		ans = max(ans, expand(A, i, i+1))
	}
	return ans
}

func LongestPallindromeLength(A string) {
	fmt.Printf("LongestPallindromeLength1 of string %s is:  %d\n", A, LongestPallindromeLength1(A))
	LongestPallindromeLength2(A)
}

func expand(A string, i, j int) int {
	ans := 0
	for i >= 0 && j < len(A) && A[i] == A[j] {
		ans = j - i + 1
		i--
		j++
	}
	return ans
}
