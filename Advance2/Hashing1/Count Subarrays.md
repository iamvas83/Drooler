## Count Subarrays

```
Problem Description

Misha likes finding all Subarrays of an Array. Now she gives you an array A of N elements and told you to find the number of subarrays of A, that have unique elements.

Since the number of subarrays could be large, return value % 109 +7.



Problem Constraints

1 <= N <= 105

1 <= A[i] <= 106



Input Format

The only argument given is an Array A, having N integers.



Output Format

Return the number of subarrays of A, that have unique elements.



Example Input

Input 1:

 A = [1, 1, 3]
Input 2:

 A = [2, 1, 2]


Example Output

Output 1:

 4
Output 1:

 5


Example Explanation

Explanation 1:

 Subarrays of A that have unique elements only:
 [1], [1], [1, 3], [3]
Explanation 2:

 Subarrays of A that have unique elements only:
 [2], [1], [2, 1], [1, 2], [2]

```
``` go
package main

import "fmt"

func countUniqueSubarrays(A []int) int {
    const MOD = 1000000007
    n := len(A)
    lastSeen := make(map[int]int) // To track the last seen index of each element
    start := 0
    result := 0

    for end := 0; end < n; end++ {
        // If the element is already in the window, move the start pointer
        if lastIdx, found := lastSeen[A[end]]; found && lastIdx >= start {
            start = lastIdx + 1
        }

        // Count subarrays ending at the current index
        result += end - start + 1
        result %= MOD

        // Update the last seen index of the current element
        lastSeen[A[end]] = end
    }

    return result
}

func main() {
    A := []int{1, 2, 1, 3, 2} // Example input
    fmt.Println(countUniqueSubarrays(A))
}
```

``` go
/**
 * @input A : Integer array
 * 
 * @Output Integer
 */
func solve(A []int )  (int) {
    mp := make(map[int]int)
    l, r, mod := 0, 0, 1000 * 1000 * 1000 + 7
    ans := 0
    // l and r represent the window which we are checking
    for i := 0; i < len(A); i++ {
        x := A[i]
        if mp[x] != 0 {
            // if we already have x in window, start removing elements from the left
            for mp[x] != 0 {
                mp[A[l]] =  0
                l++
            }
        }
        mp[x] = 1
        // number of subarrays ending at index r
        ans += r - l + 1
        ans %= mod
        r++
    }
    return ans
}
```