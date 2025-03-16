## Make String Palindrome
```
Problem Description

Given a string A of size N consisting only of lowercase alphabets. The only operation allowed is to insert characters in the beginning of the string.


Find and return how many minimum characters are needed to be inserted to make the string a palindrome string.




Problem Constraints

1 <= N <= 106



Input Format

The only argument given is a string A.



Output Format

Return an integer denoting the minimum characters needed to be inserted in the beginning to make the string a palindrome string.



Example Input

Input 1:

 A = "abc"
Input 2:

 A = "bb"


Example Output

Output 1:

 2
Output 2:

 0


Example Explanation

Explanation 1:

 Insert 'b' at beginning, string becomes: "babc".
 Insert 'c' at beginning, string becomes: "cbabc".
Explanation 2:

 There is no need to insert any character at the beginning as the string is already a palindrome. 
```

``` go 
/**
 * @input A : String
 * 
 * @Output Integer
 */
func solve(A string )  (int) {
    rev:=reverseStr(A)
    concatStr:=A+"#"+rev
    lps := computeLPS(concatStr)

	// Minimum insertions required
	return len(A) - lps[len(lps)-1]
}

func computeLPS(A string)[]int{
    length:=0
    lps:=make([]int,len(A))
    for i:=1;i<len(A);i++{

        for length>0 && A[i]!=A[length]{
            length=lps[length-1]
        }

        if A[i]==A[length]{
            length++
        }

        lps[i]=length
    }
    return lps
}
func reverseStr(A string) string{
    r:=[]rune(A)
    j:=len(r)-1
    i:=0

    for i<j{
        r[i],r[j]=r[j],r[i]
        i++
        j--
    }

    return string(r)
}
```