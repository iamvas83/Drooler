~~~
Problem Description

Given an array of integers A, sort the array into a wave-like array and return it.
In other words, arrange the elements into a sequence such that

a1 >= a2 <= a3 >= a4 <= a5..... 
NOTE: If multiple answers are possible, return the lexicographically smallest one.



Problem Constraints

1 <= len(A) <= 106
0 <= A[i] <= 106



Input Format

The first argument is an integer array A.



Output Format

Return an array arranged in the sequence as described.



Example Input

Input 1:

A = [1, 2, 3, 4]
Input 2:

A = [1, 2]


Example Output

Output 1:

[2, 1, 4, 3]
Output 2:

[2, 1]


Example Explanation

Explanation 1:

One possible answer : [2, 1, 4, 3]
Another possible answer : [4, 1, 3, 2]
First answer is lexicographically smallest. So, return [2, 1, 4, 3].
Explanation 1:

Only possible answer is [2, 1].
~~~
~~~ go
/**
 * @input A : Integer array
 * 
 * @Output Integer array.
 */
import "sort"
func wave(A []int )  ([]int) {
    sort.Ints(A)
    for i:=0;i<len(A)-1;i+=2{
        swap(A,i,i+1)
    }
    return A
}

func swap(A []int, x int, y int)[]int{
    temp:=A[x]
    A[x]=A[y]
    A[y]=temp
    return A
}
~~~