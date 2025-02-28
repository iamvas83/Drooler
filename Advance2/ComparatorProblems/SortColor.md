## Sort by Color

```
Problem Description

Given an array with N objects colored red, white, or blue, sort them so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will represent the colors as,

red -> 0
white -> 1
blue -> 2

Note: Using the library sort function is not allowed.



Problem Constraints

1 <= N <= 1000000
0 <= A[i] <= 2


Input Format

First and only argument of input contains an integer array A.


Output Format

Return an integer array in asked order


Example Input

Input 1 :
    A = [0, 1, 2, 0, 1, 2]
Input 2:

    A = [0]

```
``` go
/**
 * @input A : Integer array
 * 
 * @Output Integer array.
 */

func sortColors(A []int )  ([]int) {
low,mid,high:=0,0,len(A)-1

    for mid<=high{

        if A[mid]==0{
            A[low],A[mid]=A[mid],A[low]
            low++
            mid++
        }else if A[mid]==1{
            mid++
        }else{ //mid ==2
            A[mid],A[high]=A[high],A[mid]
            high--
        }
    }
    return A
}

```