## Pairs with Given Difference
```
Problem Description

Given an one-dimensional integer array A of size N and an integer B.

Count all distinct pairs with difference equal to B.

Here a pair is defined as an integer pair (x, y), where x and y are both numbers in the array and their absolute difference is B.



Problem Constraints

1 <= N <= 104

0 <= A[i], B <= 105



Input Format

First argument is an one-dimensional integer array A of size N.

Second argument is an integer B.



Output Format

Return an integer denoting the count of all distinct pairs with difference equal to B.



Example Input

Input 1:

 A = [1, 5, 3, 4, 2]
 B = 3
Input 2:

 A = [8, 12, 16, 4, 0, 20]
 B = 4
Input 3:

 A = [1, 1, 1, 2, 2]
 B = 0


Example Output

Output 1:

 2
Output 2:

 5
Output 3:

 2


Example Explanation

Explanation 1:

 There are 2 unique pairs with difference 3, the pairs are {1, 4} and {5, 2} 
Explanation 2:

 There are 5 unique pairs with difference 4, the pairs are {0, 4}, {4, 8}, {8, 12}, {12, 16} and {16, 20} 
Explanation 3:

 There are 2 unique pairs with difference 0, the pairs are {1, 1} and {2, 2}.
```
``` go

import "sort"
func solve(A []int, sum int) int {
i, j := 0, 1
	cnt := 0
	sort.Ints(A)
	n := len(A)
	for j < n {
		if A[j]-A[i] == sum {
			if i != j {
				cnt++

				i++
				j++

				for i < n && A[i] == A[i-1] {
					i++
				}

				for j < n && A[j] == A[j-1] {
					j++
				}
			} 
		} else if A[j]-A[i] < sum {
			j++
		} else {
			for A[j]-A[i] > sum {
				i++
			}
		}
		if i == j {
			j++
		}
	}
	return cnt
}
```