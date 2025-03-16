## Rabin karp
```
Rabin-Karp Algorithm for Pattern Matching
The Rabin-Karp algorithm is an efficient string searching algorithm that uses hashing to find the occurrence of a pattern within a text. It is particularly useful when searching for multiple patterns in the same text.

Key Idea
Instead of directly comparing the pattern with every substring in the text, we compare their hash values.
If the hash values match, only then do we perform a character-by-character comparison to confirm the match.
This reduces the number of unnecessary comparisons, making the algorithm efficient.
Steps in Rabin-Karp Algorithm
Calculate Hash for the Pattern:
Compute the hash value for the pattern using a suitable hash function.

Calculate Hash for the First Window in the Text:
The window size should be the same as the pattern's length.

Slide the Window Over the Text:

For each position, compare the hash value of the current window with the pattern's hash.
If they match, perform a character-by-character comparison to confirm the match.
Use the rolling hash technique to efficiently compute the next window's hash value.
Rolling Hash Formula:
To calculate the new hash value when sliding the window by one position:

new_hash=(old_hash−leading_char×d^m−1)×d+next_char

Where:

d = Base value (commonly 256 for ASCII characters)
m = Length of the pattern
q = A prime number to minimize collisions
```

~~~ go
package main

import (
	"fmt"
)

const d = 256            // Base for hashing (ASCII size)
const q = 101            // Prime number to reduce collisions

// Rabin-Karp Algorithm for Pattern Matching
func rabinKarp(text, pattern string) []int {
	n := len(text)
	m := len(pattern)
	p := 0  // Hash value for pattern
	t := 0  // Hash value for text window
	h := 1  // Value of d^(m-1)
	result := []int{}

	// Compute h = d^(m-1) % q
	for i := 0; i < m-1; i++ {
		h = (h * d) % q
	}

	// Calculate initial hash values for pattern and first text window
	for i := 0; i < m; i++ {
		p = (d*p + int(pattern[i])) % q
		t = (d*t + int(text[i])) % q
	}

	// Slide the window over the text
	for i := 0; i <= n-m; i++ {
		// If hash values match, check characters one by one
		if p == t {
			if text[i:i+m] == pattern {
				result = append(result, i)
			}
		}

		// Compute hash for next window
		if i < n-m {
			t = (d*(t-int(text[i])*h) + int(text[i+m])) % q

			// Handle negative hash values (Go's % operator may return negative)
			if t < 0 {
				t += q
			}
		}
	}

	return result
}

func main() {
	text := "ababcababcabc"
	pattern := "abc"
	result := rabinKarp(text, pattern)

	if len(result) == 0 {
		fmt.Println("Pattern not found")
	} else {
		fmt.Println("Pattern found at indices:", result)
	}
}
~~~

~~~
Explanation
✅ d = 256: ASCII character set size (common for text).
✅ q = 101: Prime number for hash collision reduction.
✅ Hash Calculation: Both pattern and text window hashes are calculated.
✅ Sliding Window:

If the hash values match, a direct string comparison is performed to confirm the match.
The rolling hash formula efficiently calculates the next hash value.
Sample Output

Pattern found at indices: [2 7 10]

Complexity Analysis

Time Complexity: O(n + m) on average (due to efficient hash calculations).
Worst-Case Complexity: O(n * m) — Occurs when multiple hash collisions happen.
Space Complexity: O(1) — Only a few variables are used.

Advantages
✅ Efficient for multiple pattern searches.
✅ Excellent performance for large texts.
✅ Reduces unnecessary character comparisons using hash values.

Disadvantages
❗️ May suffer from hash collisions (mitigated by using a good prime number like q = 101).
❗️ Extra computation is required for hash value maintenance.

Best Use Cases
✅ Searching for multiple patterns in the same text.
✅ Detecting plagiarism or substring searches in large documents.
✅ Efficient string matching in bioinformatics, logs, and databases.

~~~