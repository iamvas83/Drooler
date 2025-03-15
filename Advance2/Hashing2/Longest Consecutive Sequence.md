## Longest Consecutive Sequence

```
Problem Description

Given an unsorted integer array A of size N.


Find the length of the longest set of consecutive elements from array A.




Problem Constraints

1 <= N <= 106

-106 <= A[i] <= 106



Input Format

First argument is an integer array A of size N.



Output Format

Return an integer denoting the length of the longest set of consecutive elements from the array A.



Example Input

Input 1:

A = [100, 4, 200, 1, 3, 2]
Input 2:

A = [2, 1]


Example Output

Output 1:

 4
Output 2:

 2


Example Explanation

Explanation 1:

 The set of consecutive elements will be [1, 2, 3, 4].
Explanation 2:

 The set of consecutive elements will be [1, 2].
```
``` go
package main

import "fmt"

func longestConsecutive(A []int) int {
    if len(A) == 0 {
        return 0
    }

    // Step 1: Insert all elements into a set
    numSet := make(map[int]bool)
    for _, num := range A {
        numSet[num] = true
    }

    maxLength := 0

    // Step 2: Find the longest sequence
    for num := range numSet {
        // Check if this is the start of a sequence
        if !numSet[num-1] {
            currentNum := num
            length := 1

            // Count the length of the sequence
            for numSet[currentNum+1] {
                currentNum++
                length++
            }

            // Update the maximum length
            if length > maxLength {
                maxLength = length
            }
        }
    }

    return maxLength
}

func main() {
    A := []int{100, 4, 200, 1, 3, 2} // Example input
    fmt.Println(longestConsecutive(A)) // Output: 4
}
```