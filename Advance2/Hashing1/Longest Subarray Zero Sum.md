##  Longest Subarray Zero Sum

``` 
Problem Description

Given an array A of N integers.
Find the length of the longest subarray in the array which sums to zero.

If there is no subarray which sums to zero then return 0.



Problem Constraints

1 <= N <= 105
-109 <= A[i] <= 109


Input Format

Single argument which is an integer array A.


Output Format

Return an integer.


Example Input

Input 1:

 A = [1, -2, 1, 2]
Input 2:

 A = [3, 2, -1]


Example Output

Output 1:

3
Output 2:

0


Example Explanation

Explanation 1:

 [1, -2, 1] is the largest subarray which sums up to 0.
Explanation 2:

 No subarray sums up to 0.
```
``` go
/**
 * @input A : Integer array
 * 
 * @Output Integer
 */
func solve(A []int )  (int) {
m := make(map[int]int)
	var ans int
	pf := 0
	for i := 1; i <= len(A); i++ {
		pf += A[i-1]
		if _, ok := m[pf]; ok {
			ans = max(ans, i-m[pf])
		} else {
			m[pf] = 1
		}

	}
	return ans

}

func max(A,B int)int{
    if A>B{
        return A
    }else{
        return B
    }
}


package main

import "fmt"

func longestZeroSumSubarray(A []int) int {
    n := len(A)
    currentSum := 0
    maxLength := 0
    sumIndexMap := make(map[int]int) // Hash map to store the first occurrence of each sum

    for i := 0; i < n; i++ {
        currentSum += A[i]

        // If the sum is zero, the subarray from the start to i sums to zero
        if currentSum == 0 {
            maxLength = i + 1
        }

        // If the sum has been seen before, calculate the subarray length
        if prevIndex, found := sumIndexMap[currentSum]; found {
            maxLength = max(maxLength, i-prevIndex)
        } else {
            // Store the first occurrence of the sum
            sumIndexMap[currentSum] = i
        }
    }

    return maxLength
}

// Utility function to calculate the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    A := []int{1, 2, -3, 3, 1, -4, 2, 2} // Example input
    fmt.Println(longestZeroSumSubarray(A))
}

package main

import "fmt"

func longestZeroSumSubarray(A []int) int {
    n := len(A)
    prefixSum := 0
    maxLength := 0
    prefixMap := make(map[int]int) // To store the first occurrence of a prefix sum

    for i := 0; i < n; i++ {
        prefixSum += A[i]

        // If prefix sum is zero, the subarray from 0 to i sums to zero
        if prefixSum == 0 {
            maxLength = i + 1
        }

        // If prefix sum is seen before, calculate the subarray length
        if prevIndex, found := prefixMap[prefixSum]; found {
            maxLength = max(maxLength, i-prevIndex)
        } else {
            // Store the first occurrence of the prefix sum
            prefixMap[prefixSum] = i
        }
    }

    return maxLength
}

// Utility function to calculate the maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    A := []int{1, 2, -3, 3, 1, -4, 2, 2} // Example input
    fmt.Println(longestZeroSumSubarray(A))
}
```