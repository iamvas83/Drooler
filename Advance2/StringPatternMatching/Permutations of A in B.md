## Permutations of A in B

```
Problem Description

You are given two strings, A and B, of size N and M, respectively.

You have to find the count of all permutations of A present in B as a substring. You can assume a string will have only lowercase letters.




Problem Constraints

1 <= N < M <= 105



Input Format

Given two arguments, A and B of type String.



Output Format

Return a single integer, i.e., number of permutations of A present in B as a substring.



Example Input

Input 1:

 A = "abc"
 B = "abcbacabc"
Input 2:

 A = "aca"
 B = "acaa"





Example Output

Output 1:

 5
Output 2:

 2


Example Explanation

Explanation 1:

 Permutations of A that are present in B as substring are:
    1. abc
    2. cba
    3. bac
    4. cab
    5. abc
    So ans is 5.
Explanation 2:

 Permutations of A that are present in B as substring are:
    1. aca
    2. caa 
```

``` go
/**
 * @input A : String
 * @input B : String
 * 
 * @Output Integer
 */
import "reflect"
func solve(B string , A string )  (int) {
    n:=len(A)
    m:=len(B)
    ans:=0
    B_freq:=make([]int,26)
    A_freq:=make([]int,26)
    
    for i:=0;i<m;i++{
        B_freq[B[i]-'a']++
    }

    for i:=0;i<m;i++{
        A_freq[A[i]-'a']++
    }

    if reflect.DeepEqual(B_freq,A_freq){
        ans++
    }

    s:=1
    e:=m

    for e<n{
        out:=A[s-1]
        in:=A[e]

        A_freq[out-'a']--
        A_freq[in-'a']++
        if reflect.DeepEqual(A_freq,B_freq){
            ans++
        }
        s++
        e++
    }
    return ans

}

// func checkIfEqual(A,B []int)bool{
//     flag :=false
//     for i:=0;i<len(A);i++{
//         if A[i]==B[i]{
//             flag=true
//         }else{
//             flag=false
//             break
//         }
//     }
//     return flag
// }
```