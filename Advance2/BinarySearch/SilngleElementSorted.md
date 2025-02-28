## Single Element in Sorted Array

~~~
Problem Description

Given a sorted array of integers A where every element appears twice except for one element which appears once, find and return this single element that appears only once.

Elements which are appearing twice are adjacent to each other.

NOTE: Users are expected to solve this in O(log(N)) time.



Problem Constraints

1 <= |A| <= 100000

1 <= A[i] <= 10^9



Input Format

The only argument given is the integer array A.



Output Format

Return the single element that appears only once.



Example Input

Input 1:

A = [1, 1, 7]
Input 2:

A = [2, 3, 3]


Example Output

Output 1:

 7
Output 2:

 2


Example Explanation

Explanation 1:

 7 appears once
Explanation 2:

 2 appears once
~~~
~~~ go
/**
 * @input A : Integer array
 * 
 * @Output Integer
 */
func solve(A []int )  (int) {
    n := len(A)
    l, r := 0, n-1

    for l < r {
        mid := l + (r-l)/2
        
        // Ensure mid is even to check pairs correctly
        if mid%2 == 1 {
            mid--
        }

        // If pair starts correctly, move right
        if mid+1 < n && A[mid] == A[mid+1] {
            l = mid + 2
        } else {
            r = mid
        }
        
}
return A[l] 
}
~~~

### Sol 2

``` go
/**
 * @input A : Integer array
 * 
 * @Output Integer
 */
func solve(A []int )  (int) {
    n := len(A)
    l := 0
    r := n - 1
    var ans int
    for l<=r{
    
    if l==r{
        ans= A[l]
        break
    }

    mid:=l+(r-l)/2

    if mid%2==0{
        if A[mid]==A[mid+1]{
            l=mid+2
        }else{
            r=mid
        }
    }else if mid%2==1{
        if A[mid]==A[mid-1]{
            l=mid+1
        }else{
            r=mid-1
        }
    }
    }
    return ans
}
```
