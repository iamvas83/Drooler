## Distinct Numbers in Window

```
Problem Description

You are given an array of N integers, A1, A2 ,..., AN and an integer B. Return the of count of distinct numbers in all windows of size B.

Formally, return an array of size N-B+1 where i'th element in this array contains number of distinct elements in sequence Ai, Ai+1 ,..., Ai+B-1.

NOTE: if B > N, return an empty array.



Problem Constraints

1 <= N <= 106

1 <= A[i] <= 109


Input Format

First argument is an integer array A
Second argument is an integer B.



Output Format

Return an integer array.



Example Input

Input 1:

 A = [1, 2, 1, 3, 4, 3]
 B = 3
Input 2:

 A = [1, 1, 2, 2]
 B = 1


Example Output

Output 1:

 [2, 3, 3, 2]
Output 2:

 [1, 1, 1, 1]

```

``` go
/**
 * @input A : Integer array
 * @input B : Integer
 * 
 * @Output Integer array.
 */
func dNums(A []int , k int )  ([]int) {
    m := make(map[int]int)
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		if _, ok := m[A[i]]; ok {
			m[A[i]] = m[A[i]] + 1
		} else {
			m[A[i]] = 1
		}
	}
	ans = append(ans, len(m))

	s := 1
	e := k

	for e < len(A) {
		out := A[s-1]
		in := A[e]
		if _, ok := m[out]; ok {
			m[out] = m[out] - 1
		}
		if _, ok := m[in]; ok {
			m[in] = m[in] + 1
		} else {
			m[in] = 1
		}

		if m[out] == 0 {
			delete(m, out)
		}
		ans = append(ans, len(m))
		s++
		e++

	}
	return ans
}
```