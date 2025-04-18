## Allocate Books

```
Problem Description

Given an array of integers A of size N and an integer B.

The College library has N books. The ith book has A[i] number of pages.

You have to allocate books to B number of students so that the maximum number of pages allocated to a student is minimum.

A book will be allocated to exactly one student.
Each student has to be allocated at least one book.
Allotment should be in contiguous order, for example: A student cannot be allocated book 1 and book 3, skipping book 2.
Calculate and return that minimum possible number.

NOTE: Return -1 if a valid assignment is not possible.




Problem Constraints

1 <= N <= 105
1 <= A[i], B <= 105



Input Format

The first argument given is the integer array A.
The second argument given is the integer B.



Output Format

Return that minimum possible number.



Example Input

Input 1:
A = [12, 34, 67, 90]
B = 2
Input 2:
A = [12, 15, 78] 
B = 4


Example Output

Output 1:
113
Output 2:
-1


Example Explanation

Explanation 1:

There are two students. Books can be distributed in following fashion : 
1)  [12] and [34, 67, 90]
    Max number of pages is allocated to student 2 with 34 + 67 + 90 = 191 pages
2)  [12, 34] and [67, 90]
    Max number of pages is allocated to student 2 with 67 + 90 = 157 pages 
3)  [12, 34, 67] and [90]
    Max number of pages is allocated to student 1 with 12 + 34 + 67 = 113 pages
    Of the 3 cases, Option 3 has the minimum pages = 113.
Explanation 2:
Each student has to be allocated at least one book.  
But the Total number of books is less than the number of students.
Thus each student cannot be allocated to atleast one book.

Therefore, the result is -1.
```

``` go
/**
 * @input A : Integer array
 * @input B : Integer
 * 
 * @Output Integer
 */
import "math"

func books(A []int , B int )  (int) {
    if len(A)<B{return -1}
    max:=math.MinInt
    for i:=0;i<len(A);i++{
        if A[i]>max{
            max=A[i]

        }
    }

    var sum int
    for i:=0;i<len(A);i++{
        sum+=A[i]
    }

    l:=max
    r:=sum
    ans:=0
    for l<=r{
        mid:=l+(r-l)/2
        if isfeasible(A,B,mid){
            ans=mid
            r=mid-1
        }else{
            l=mid+1
        }

    }
    return ans
}

func isfeasible(A []int, k int, ans int )bool{
    req:=1
    sum:=0

    for i:=0;i<len(A);i++{
        if sum+A[i]>ans{
            req++
            sum=A[i]
        }else{
            sum+=A[i]

        }
    }
    return req<=k
}

```