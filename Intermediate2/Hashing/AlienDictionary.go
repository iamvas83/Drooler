package hashing

import (
	"fmt"
)

// Function to check if the words are sorted according to the alien language order
func AreWordsSorted(words []string, order string) int {
	// Create a map to store the order of each character
	charOrder := make(map[rune]int)
	for i, char := range order {
		charOrder[char] = i
	}

	// Compare each word with the next one
	for i := 0; i < len(words)-1; i++ {
		if !IsSorted(words[i], words[i+1], charOrder) {
			return 0
		}
	}

	return 1
}

// Helper function to compare two words based on the given character order
func IsSorted(word1, word2 string, charOrder map[rune]int) bool {
	len1, len2 := len(word1), len(word2)
	minLen := len1
	if len2 < len1 {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		char1, char2 := rune(word1[i]), rune(word2[i])
		if charOrder[char1] < charOrder[char2] {
			return true
		} else if charOrder[char1] > charOrder[char2] {
			return false
		}
	}

	// If one word is a prefix of the other, the shorter word should come first
	return len1 <= len2
}

func Test() {
	// Example Input 1
	words1 := []string{"hello", "scaler", "interviewbit"}
	order1 := "adhbcfegskjlponmirqtxwuvzy"
	fmt.Println(AreWordsSorted(words1, order1)) // Output: 1

	// Example Input 2
	words2 := []string{"fine", "none", "bush"}
	order2 := "qwertyuiopasdfghjklzxcvbnm"
	fmt.Println(AreWordsSorted(words2, order2)) // Output: 0
}
