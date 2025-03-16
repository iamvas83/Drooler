## Isomorphic Strings

```
Problem Description

Given two strings A and B, determine if they are isomorphic.


A and B are called isomorphic strings if all occurrences of each character in A can be replaced with another character to get B and vice-versa.




Problem Constraints

1 <= |A| <= 100000

1 <= |B| <= 100000

A and B contain only lowercase characters.



Input Format

The first Argument is string A.


The second Argument is string B.




Output Format

Return 1 if strings are isomorphic, 0 otherwise.



Example Input

Input 1:

A = "aba"
B = "xyx"
Input 2:

A = "cvv"
B = "xyx"


Example Output

Output 1:

 1
Output 2:

 0


Example Explanation

Explanation 1:

 Replace 'a' with 'x', 'b' with 'y'.
Explanation 2:

 A cannot be converted to B.
```
``` go
/**
 * @input A : String
 * @input B : String
 * 
 * @Output Integer
 */
func solve(A string , B string )  (int) {

    if len(A) != len(B) {
		return 0
	}

	// Step 2: Mapping logic
	mapAtoB := make(map[rune]rune)
	mapBtoA := make(map[rune]rune)

	for i := 0; i < len(A); i++ {
		charA, charB := rune(A[i]), rune(B[i])

		// Step 3: Check mappings
		if mappedB, exists := mapAtoB[charA]; exists && mappedB != charB {
			return 0
		}
		if mappedA, exists := mapBtoA[charB]; exists && mappedA != charA {
			return 0
		}

		// Step 4: Add mappings
		mapAtoB[charA] = charB
		mapBtoA[charB] = charA
	}

	return 1
}
```