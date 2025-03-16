## Rabin Karp2

```
Problem Description

Given two string A and B of length N and M respectively consisting of lowercase letters. Find the number of occurrences of B in A.


Problem Constraints

1 <= M <= N <= 105



Input Format

Two argument A and B are strings


Output Format

Return an integer denoting the number of occurrences of B in A


Example Input

Input 1:
A = "acbac"
B = "ac"
Input 2:
A = "aaaa"
B = "aa"


Example Output

Output 1:
2
Output 2:
3


Example Explanation

For Input 1:
The string "ac" occurs twice in "acbac".
For Input 2:
The string "aa" occurs thrice in "aaaa".
```
``` go
/**
 * @input A : String
 * @input B : String
 * 
 * @Output Integer
 */
func solve(A string , patt string )  (int) {
    n := len(A)
	k := len(patt)
	d := 256
	q := 101
	h := 1
	ans := 0
	for i := 0; i < k; i++ {
		h = (h * d) % q
	}
	p := 0
	t := 0

	for i := 0; i < k; i++ {
		p = (p*d + int(patt[i])) % q
		t = (t*d + int(A[i])) % q
	}

	for i := 0; i < n-k+1; i++ {
		flag := true
		for j := 0; j < k; j++ {
			if A[i+j] != patt[j] {
				flag = false
				break
			}
		}
		if flag {
			ans++
		}
		if i < n-k {
			t = (d*(t-int(A[i])*h) + int(A[i+k])) % q
		}
		if t < 0 {
			t = t + q
		}
	}
	return ans
}


```