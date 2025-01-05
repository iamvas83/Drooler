## Toggle Case ##

Problem Description

You are given a character string A having length N, consisting of only lowercase and uppercase latin letters.

You have to toggle case of each character of string A. For e.g 'A' is changed to 'a', 'e' is changed to 'E', etc.



Problem Constraints

1 <= N <= 105

A[i] ∈ ['a'-'z', 'A'-'Z']



Input Format

First and only argument is a character string A.



Output Format

Return a character string.



Example Input

Input 1:

 A = "Hello" 
Input 2:

 A = "tHiSiSaStRiNg" 


Example Output

Output 1:

 hELLO 
Output 2:

 ThIsIsAsTrInG 


Example Explanation

Explanation 1:

 'H' changes to 'h'
 'e' changes to 'E'
 'l' changes to 'L'
 'l' changes to 'L'
 'o' changes to 'O'
Explanation 2:

 "tHiSiSaStRiNg" changes to "ThIsIsAsTrInG".

 ## CountSort (Sort Lower case)

 Problem Description

Given an array A. Sort this array using Count Sort Algorithm and return the sorted array.


Problem Constraints

1 <= |A| <= 105
1 <= A[i] <= 105


Input Format

The first argument is an integer array A.


Output Format

Return an integer array that is the sorted array A.


Example Input

Input 1:
A = [1, 3, 1]
Input 2:
A = [4, 2, 1, 3]


Example Output

Output 1:
[1, 1, 3]
Output 2:
[1, 2, 3, 4]


Example Explanation

For Input 1:
The array in sorted order is [1, 1, 3].
For Input 2:
The array in sorted order is [1, 2, 3, 4].

## Simple Reverse

Problem Description

Given an array A. Sort this array using Count Sort Algorithm and return the sorted array.


Problem Constraints

1 <= |A| <= 105
1 <= A[i] <= 105


Input Format

The first argument is an integer array A.


Output Format

Return an integer array that is the sorted array A.


Example Input

Input 1:
A = [1, 3, 1]
Input 2:
A = [4, 2, 1, 3]


Example Output

Output 1:
[1, 1, 3]
Output 2:
[1, 2, 3, 4]


Example Explanation

For Input 1:
The array in sorted order is [1, 1, 3].
For Input 2:
The array in sorted order is [1, 2, 3, 4].

``` go

func solve(A string )  (string) {
    runes:=[]rune(A)
    for i,j:=0,len(runes)-1;i<j;i,j=i+1,j-1{
        runes[i],runes[j]=runes[j],runes[i]
    }
    return string(runes)
}

```

## Reverse the String

Problem Description

You are given a string A of size N.


Return the string A after reversing the string word by word.

NOTE:

A sequence of non-space characters constitutes a word.
Your reversed string should not contain leading or trailing spaces, even if it is present in the input string.
If there are multiple spaces between words, reduce them to a single space in the reversed string.



Problem Constraints

1 <= N <= 3 * 105



Input Format

The only argument given is string A.



Output Format

Return the string A after reversing the string word by word.



Example Input

Input 1:
A = "the sky is blue"
Input 2:
A = "this is ib"


Example Output

Output 1:
"blue is sky the"
Output 2:
"ib is this"    


Example Explanation

Explanation 1:
We reverse the string word by word so the string becomes "blue is sky the".
Explanation 2:
We reverse the string word by word so the string becomes "ib is this".


``` go

import "strings"
func solve(A string )  (string) {
    words:=strings.Fields(A)
    for i,j:=0,len(words)-1;i<j;i,j=i+1,j-1{
        words[i],words[j]=words[j],words[i]
    }
    return strings.Join(words, " ")
}

```

## Longest Palindromic Substring

Problem Description

Given a string A of size N, find and return the longest palindromic substring in A.

Substring of string A is A[i...j] where 0 <= i <= j < len(A)

Palindrome string:
A string which reads the same backwards. More formally, A is palindrome if reverse(A) = A.

Incase of conflict, return the substring which occurs first ( with the least starting index).



Problem Constraints

1 <= N <= 6000



Input Format

First and only argument is a string A.



Output Format

Return a string denoting the longest palindromic substring of string A.



Example Input

Input 1:
A = "aaaabaaa"
Input 2:
A = "abba


Example Output

Output 1:
"aaabaaa"
Output 2:
"abba"


Example Explanation

Explanation 1:
We can see that longest palindromic substring is of length 7 and the string is "aaabaaa".
Explanation 2:
We can see that longest palindromic substring is of length 4 and the string is "abba".


##  Longest Common Prefix

Problem Description

Given the array of strings A, you need to find the longest string S, which is the prefix of ALL the strings in the array.


The longest common prefix for a pair of strings S1 and S2 is the longest string S which is the prefix of both S1 and S2.

Example: the longest common prefix of "abcdefgh" and "abcefgh" is "abc".




Problem Constraints

0 <= sum of length of all strings <= 1000000



Input Format

The only argument given is an array of strings A.



Output Format

Return the longest common prefix of all strings in A.



Example Input

Input 1:


A = ["abcdefgh", "aefghijk", "abcefgh"]
Input 2:

A = ["abab", "ab", "abcd"];






Example Output

Output 1:

"a"
Output 2:

"ab"


Example Explanation

Explanation 1:

Longest common prefix of all the strings is "a".
Explanation 2:

Longest common prefix of all the strings is "ab".

``` c++

#include <iostream>
#include <vector>
#include <string>
using namespace std;

string longestCommonPrefix(vector<string>& strs) {
    if (strs.empty()) return "";
    string prefix = strs[0];
    for (int i = 1; i < strs.size(); i++) {
        while (strs[i].find(prefix) != 0) {
            prefix = prefix.substr(0, prefix.size() - 1);
            if (prefix.empty()) return "";
        }
    }
    return prefix;
}

int main() {
    vector<string> strs = {"flower", "flow", "flight"};
    cout << "Longest Common Prefix: " << longestCommonPrefix(strs) << endl;
    return 0;
}

```

``` python

def longest_common_prefix(strs):
    if not strs:
        return ""
    
    prefix = strs[0]
    for s in strs[1:]:
        while not s.startswith(prefix):
            prefix = prefix[:-1]
            if not prefix:
                return ""
    return prefix

# Test case
strs = ["flower", "flow", "flight"]
print("Longest Common Prefix:", longest_common_prefix(strs))

```

``` go

package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, str := range strs[1:] {
		for strings.Index(str, prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println("Longest Common Prefix:", longestCommonPrefix(strs))
}

```

## Check anagrams

Problem Description

You are given two lowercase strings A and B each of length N. Return 1 if they are anagrams to each other and 0 if not.

Note : Two strings A and B are called anagrams to each other if A can be formed after rearranging the letters of B.


Problem Constraints

1 <= N <= 105
A and B are lowercase strings


Input Format

Both arguments A and B are a string.


Output Format

Return 1 if they are anagrams and 0 if not


Example Input

Input 1:
A = "cat"
B = "bat"
Input 2:
A = "secure"
B = "rescue"


Example Output

Output 1:
0
Output 2:
1


Example Explanation

For Input 1:
The words cannot be rearranged to form the same word. So, they are not anagrams.
For Input 2:
They are an anagram.

``` go
import "sort"
func solve(A string , B string )  (int) {
    runes := []rune(A)

	sort.Slice(runes, func(i int, j int) bool {
		return runes[i] < runes[j]
	})
	runes2 := []rune(B)

	sort.Slice(runes2, func(i int, j int) bool {
		return runes2[i] < runes2[j]
	})
	if string(runes) == string(runes2) {
		return 1
	} else {
		return 0
	}

}

```



## Code explanation for Longest Prefix Sum

Problem Description
Given an array of strings, find the longest string P such that:

P is a prefix of every string in the array.
A prefix is the initial portion of a string (e.g., "fl" is a prefix of "flower").
For example:

Input: ["flower", "flow", "flight"]
Output: "fl" (common prefix among all strings)

Key Idea
The longest common prefix is obtained by repeatedly reducing a candidate prefix until it matches the start of every string in the array.

Start with the first string as the prefix.

Initially assume the entire first string (strs[0]) is the prefix.
This is the "candidate prefix."
Iteratively compare with other strings.

For each subsequent string in the array:
Check if the current prefix is actually a prefix of the string.
If not, shrink the prefix from the end by removing one character at a time until it matches the start of the string.
Terminate early if the prefix becomes empty.

If the prefix is reduced to an empty string, return "", as no common prefix exists.

Using strings.Index
In Go, strings.Index(string, substring) checks whether a substring exists in the string:

If strings.Index(str, prefix) == 0: The substring prefix is at the start of str.
If strings.Index(str, prefix) != 0: The substring is not at the start.
This function simplifies the process of checking whether one string starts with another.

Complexity Analysis
Time Complexity: 
𝑂(𝑆):

>S is the total number of characters in all strings.

>In the worst case, every character of every string is compared.

Space Complexity: 
𝑂(1):
Only a few variables are used; no additional space is required.