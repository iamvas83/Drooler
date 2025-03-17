## Window String

```
Problem Description

Given a string A and a string B, find the window with minimum length in A, which will contain all the characters in B in linear time complexity.
Note that when the count of a character c in B is x, then the count of c in the minimum window in A should be at least x.

Note:

If there is no such window in A that covers all characters in B, return the empty string.
If there are multiple such windows, return the first occurring minimum window ( with minimum start index and length )



Problem Constraints

1 <= size(A), size(B) <= 106



Input Format

The first argument is a string A.
The second argument is a string B.



Output Format

Return a string denoting the minimum window.



Example Input

Input 1:

 A = "ADOBECODEBANC"
 B = "ABC"
Input 2:

 A = "Aa91b"
 B = "ab"



Example Output

Output 1:

 "BANC"
Output 2:

 "a91b"


Example Explanation

Explanation 1:

 "BANC" is a substring of A which contains all characters of B.
Explanation 2:

 "a91b" is the substring of A which contains all characters of B.

```
``` go
/**
 * @input A : String
 * @input B : String
 * 
 * @Output string.
 */
import "math"
func minWindow(A string , B string )  (string) {
    if len(A) < len(B) {
		return ""
	}
	ans := math.MaxInt
	//n := len(A)
	hmb := map[byte]int{}
	hmstr := map[byte]int{}
	satisfaction := 0

	for i := 0; i < len(B); i++ {
		if _, ok := hmb[B[i]]; ok {
			hmb[B[i]] = hmb[B[i]] + 1
		} else {
			hmb[B[i]] = 1
		}
	}
	j, k := 0, -1

	for i, c := range A {
		cb := byte(c)
		if _, ok := hmstr[cb]; ok {
			hmstr[cb] = hmstr[cb] + 1
		} else {
			hmstr[cb] = 1
		}
		if hmb[cb] >= hmstr[cb] {
			satisfaction++
		}

		for satisfaction == len(B) {
			if i-j+1 < ans {
				ans = i - j + 1
				k = j
			}
			if hmb[A[j]] >= hmstr[A[j]] {
				satisfaction--
			}
			hmstr[A[j]] = hmstr[A[j]] - 1
			j++
		}

	}
	if k < 0 {
		return ""
	} else {
		return A[k : k+ans]
	}
}
```

``` go
package main

import (
	"fmt"
	"math"
)

func minWindow(A string, B string) string {
	if len(A) < len(B) {
		return ""
	}

	// Frequency map for string B
	neededChars := make(map[byte]int)
	for i := 0; i < len(B); i++ {
		neededChars[B[i]]++
	}

	// Variables to track window
	left, right := 0, 0
	count := 0
	minLen := math.MaxInt32
	minStart := 0

	// Map to track characters in the current window
	windowChars := make(map[byte]int)

	for right < len(A) {
		// Add character at `right` to window
		windowChars[A[right]]++

		// Check if this character helps satisfy the condition
		if neededChars[A[right]] > 0 && windowChars[A[right]] <= neededChars[A[right]] {
			count++
		}

		// Shrink the window from the left when all characters are found
		for count == len(B) {
			// Update minimum window details
			if right-left+1 < minLen {
				minLen = right - left + 1
				minStart = left
			}

			// Remove character at `left` from window
			windowChars[A[left]]--

			// If removing this character breaks the valid window condition
			if neededChars[A[left]] > 0 && windowChars[A[left]] < neededChars[A[left]] {
				count--
			}
			left++
		}

		right++
	}

	// Return the minimum window substring
	if minLen == math.MaxInt32 {
		return ""
	}
	return A[minStart : minStart+minLen]
}

func main() {
	A := "ADOBECODEBANC"
	B := "ABC"
	fmt.Println("Minimum window substring:", minWindow(A, B))
}

```