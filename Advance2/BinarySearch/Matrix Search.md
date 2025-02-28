
```
Problem Description

Given a matrix of integers A of size N x M and an integer B. Write an efficient algorithm that searches for integer B in matrix A.

This matrix A has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than or equal to the last integer of the previous row.
Return 1 if B is present in A, else return 0.

NOTE: Rows are numbered from top to bottom, and columns are from left to right.



Problem Constraints

1 <= N, M <= 1000
1 <= A[i][j], B <= 106
```

``` go
/**
 * @input A : 2D integer array 
 * @input B : Integer
 * 
 * @Output Integer
 */
func searchMatrix(A [][]int , B int )  (int) {

    n := len(A)
	m := len(A[0])
	var ans int
	for i := 0; i < n; i++ {
		if B >= A[i][0] && B <= A[i][m-1] {
			ans = i
			break
		}
	}

	l := 0
	r := m - 1

	for l <= r {
		mid := l + (r-l)/2

		if A[ans][mid] == B {
			return 1
		} else if B > A[ans][mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return 0
}

```

~~~
Optimized Approach:
Binary search on the first column to find the correct row.
Binary search within that row to check if B exists.

Time Complexity:
Finding the correct row → O(log n)
Binary search within that row → O(log m)
Total Complexity → O(log n + log m) = O(log(nm))

~~~
``` go
func searchMatrix(A [][]int, B int) int {
    n := len(A)
    m := len(A[0])
    
    // Step 1: Binary Search on the First Column to Find the Right Row
    top, bottom := 0, n-1
    var ans int
    
    for top <= bottom {
        mid := top + (bottom - top) / 2
        if B >= A[mid][0] && B <= A[mid][m-1] {
            ans = mid // Found the row where B can exist
            break
        } else if B < A[mid][0] {
            bottom = mid - 1
        } else {
            top = mid + 1
        }
    }
    
    // Step 2: Binary Search in the Identified Row
    l, r := 0, m-1
    for l <= r {
        mid := l + (r - l) / 2
        if A[ans][mid] == B {
            return 1 // Found B
        } else if B > A[ans][mid] {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    
    return 0 // B not found
}
```
