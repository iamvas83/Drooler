## Sort Array in given Order
```
Problem Description

Given two arrays of integers A and B, Sort A in such a way that the relative order among the elements will be the same as those are in B.
For the elements not present in B, append them at last in sorted order.

Return the array A after sorting from the above method.

NOTE: Elements of B are unique.



Problem Constraints

1 <= length of the array A <= 100000

1 <= length of the array B <= 100000

-10^9 <= A[i] <= 10^9






Input Format

The first argument given is the integer array A.

The second argument given is the integer array B.






Output Format

Return the array A after sorting as described.



Example Input

Input 1:

A = [1, 2, 3, 4, 5, 4]
B = [5, 4, 2]


Input 2:

A = [5, 17, 100, 11]
B = [1, 100]


Example Output

Output 1:

[5, 4, 4, 2, 1, 3]
Output 2:

[100, 5, 11, 17]


Example Explanation

Explanation 1:

Since 2, 4, 5, 4 of A are present in the array B.  So Maintaining the relative order of B.
Thus, [5, 4, 4, 2] and appending the remaining element (1, 3) in sorted order.

The Final array is [5, 4, 4, 2, 1, 3]
Explanation 2:

Since 100 of A are present in the array B.  So Maintaining the relative order of B.
Thus, [100] and appending the remaining element (5, 11, 17) in sorted order.

The Final array is [100, 5, 11, 17]
```
``` go
/**
 * @input A : Integer array
 * @input B : Integer array
 * 
 * @Output Integer array.
 */

import "sort"
func solve(A []int , B []int )  ([]int) {
    mb := map[int]int{}
	ans := make([]int, len(A))
	temp := make([]int, 0)
	for i := 0; i < len(A); i++ {
		if _, ok := mb[A[i]]; ok {
			mb[A[i]] = mb[A[i]] + 1
		} else {
			mb[A[i]] = 1
		}
	}
	count := 0
	for i := 0; i < len(B); i++ {
		if j, ok := mb[B[i]]; ok {
			for j > 0 {
				ans[count] = B[i]
				count++
				j--
				mb[B[i]] = mb[B[i]] - 1
			}

			if mb[B[i]] == 0 {
				delete(mb, B[i])
			}
		}
	}

	for val, i := range mb {
		for i > 0 {
			temp = append(temp, val)
			i--
		}

	}
	sort.Ints(temp)

	for i := 0; i < len(temp); i++ {
		ans[count] = temp[i]
		count++
	}

	return ans
}
```
``` go
package main

import (
	"fmt"
	"sort"
)

func relativeSort(A []int, B []int) []int {
	// Step 1: Frequency map for elements in A
	countMap := make(map[int]int)
	for _, num := range A {
		countMap[num]++
	}

	// Step 2: Result array to hold sorted elements
	var result []int

	// Step 3: Add elements from B in order
	for _, num := range B {
		if count, exists := countMap[num]; exists {
			for i := 0; i < count; i++ {
				result = append(result, num)
			}
			delete(countMap, num) // Remove processed elements
		}
	}

	// Step 4: Collect remaining elements
	var remaining []int
	for num, count := range countMap {
		for i := 0; i < count; i++ {
			remaining = append(remaining, num)
		}
	}

	// Step 5: Sort the remaining elements
	sort.Ints(remaining)

	// Step 6: Append sorted remaining elements
	result = append(result, remaining...)

	return result
}

func main() {
	A := []int{3, 1, 2, 2, 4, 3, 6, 7, 9, 2}
	B := []int{2, 3, 4, 6}
	fmt.Println("Sorted array:", relativeSort(A, B))
}
```