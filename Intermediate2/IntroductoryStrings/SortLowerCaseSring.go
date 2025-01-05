package introductorystrings

import "fmt"

func SortLowerCaseString(A string) {

	cnt := make([]byte, 26)
	for i := 0; i < len(A); i++ {
		cnt[A[i]-'a']++
	}

	res := ""
	for i := 0; i < len(cnt); i++ {
		for j := 0; j < int(cnt[i]); j++ {
			res += string(i + 'a')
		}
	}

	fmt.Printf("Sorting before %s and after %s the String A ", A, res)
}
