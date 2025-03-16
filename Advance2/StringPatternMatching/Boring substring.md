## Boring Substring

```
Problem Description

You are given a string A of lowercase English alphabets. Rearrange the characters of the given string A such that there is no boring substring in A.


A boring substring has the following properties:

Its length is 2.
Both the characters are consecutive, for example - "ab", "cd", "dc", "zy" etc.(If the first character is C then the next character can be either (C+1) or (C-1)).
Return 1 if it is possible to rearrange the letters of A such that there are no boring substrings in A else, return 0.




Problem Constraints

1 <= |A| <= 105



Input Format

The only argument given is a string A.



Output Format

Return 1 if it is possible to rearrange the letters of A such that there are no boring substrings in A else, return 0.



Example Input

Input 1:


 A = "abcd"
Input 2:

 A = "aab"





Example Output

Output 1:

 1
Output 2:

 0


Example Explanation

Explanation 1:

 String A can be rearranged into "cadb" or "bdac" 
Explanation 2:

 No arrangement of string A can make it free of boring substrings.

```
``` go
/**
 * @input A : String
 * 
 * @Output Integer
 */
//import "math"
import "sort"
//import "strings"
func solve(s string) int {
	oddChars := []rune{}
	evenChars := []rune{}

	// Split characters into odd and even groups
	for _, ch := range s {
		if ch%2 == 0 {
			evenChars = append(evenChars, ch)
		} else {
			oddChars = append(oddChars, ch)
		}
	}

	// Sort both groups
	sort.Slice(oddChars, func(i, j int) bool { return oddChars[i] < oddChars[j] })
	sort.Slice(evenChars, func(i, j int) bool { return evenChars[i] < evenChars[j] })

	// Check combinations: Odd + Even and Even + Odd
	if isValidCombination(oddChars, evenChars) || isValidCombination(evenChars, oddChars) {
		return 1
	}

	return 0
}

// Function to check if merging two lists avoids boring substrings
func isValidCombination(first, second []rune) bool {
	// Combine both groups
	combined := append(first, second...)

	// Check for boring pairs
	for i := 1; i < len(combined); i++ {
		if abs(int(combined[i])-int(combined[i-1])) == 1 {
			return false
		}
	}
	return true
}

// Utility function to calculate absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
