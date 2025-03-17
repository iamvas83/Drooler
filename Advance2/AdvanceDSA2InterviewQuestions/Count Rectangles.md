## Count Rectangles

```
Problem Description

Given two arrays of integers A and B of size N each, where each pair (A[i], B[i]) for 0 <= i < N represents a unique point (x, y) in a 2-D Cartesian plane.

Find and return the number of unordered quadruplet (i, j, k, l) such that (A[i], B[i]), (A[j], B[j]), (A[k], B[k]) and (A[l], B[l]) form a rectangle with the rectangle having all the sides parallel to either x-axis or y-axis.



Problem Constraints

1 <= N <= 2000
0 <= A[i], B[i] <= 109



Input Format

The first argument given is the integer array A.
The second argument given is the integer array B.



Output Format

Return the number of unordered quadruplets that form a rectangle.



Example Input

Input 1:

 A = [1, 1, 2, 2]
 B = [1, 2, 1, 2]
Input 1:

 A = [1, 1, 2, 2, 3, 3]
 B = [1, 2, 1, 2, 1, 2]



Example Output

Output 1:

 1
Output 2:

 3


Example Explanation

Explanation 1:

 All four given points make a rectangle. So, the answer is 1.
Explanation 2:

 3 quadruplets which make a rectangle are: (1, 1), (2, 1), (2, 2), (1, 2)
                                           (1, 1), (3, 1), (3, 2), (1, 2)
                                           (2, 1), (3, 1), (3, 2), (2, 2)
```
``` go
/**
 * @input A : Integer array
 * @input B : Integer array
 * 
 * @Output Integer
 */
func solve(A []int , B []int )  (int) {
// Maps to track points
	points := make(map[int]map[int]bool)

	// Populate the points map
	for i := 0; i < len(A); i++ {
		x, y := A[i], B[i]
		if points[x] == nil {
			points[x] = make(map[int]bool)
		}
		points[x][y] = true
	}

	count := 0

	// Iterate through pairs of points
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			x1, y1 := A[i], B[i]
			x2, y2 := A[j], B[j]

			// Check for valid rectangle points
			if x1 != x2 && y1 != y2 {
				if points[x1][y2] && points[x2][y1] {
					count++
				}
			}
		}
	}

	// Each rectangle is counted twice in the above logic
	return count / 2
}
```