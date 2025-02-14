## Merge Two Sorted Arrays

```
Problem Description

Given two sorted integer arrays A and B, merge B and A as one sorted array and return it as an output.

Note: A linear time complexity is expected and you should avoid use of any library function.


Problem Constraints

-2×109 <= A[i], B[i] <= 2×109
1 <= |A|, |B| <= 5×104


Input Format

First Argument is a 1-D array representing  A.
Second Argument is also a 1-D array representing B.


Output Format

Return a 1-D vector which you got after merging A and B.


Example Input

Input 1:

A = [4, 7, 9]
B = [2, 11, 19]
Input 2:

A = [1]
B = [2]


Example Output

Output 1:

[2, 4, 7, 9, 11, 19]
Output 2:

[1, 2]


Example Explanation

Explanation 1:

Merging A and B produces the output as described above.
Explanation 2:

 Merging A and B produces the output as described above.
```

## Unique Elements

```
Problem Description

You are given an array A of N elements. You have to make all elements unique. To do so, in one step you can increase any number by one.

Find the minimum number of steps.

Problem Constraints

1 <= N <= 105
1 <= A[i] <= 109


Input Format

The only argument given is an Array A, having N integers.


Output Format

Return the minimum number of steps required to make all elements unique.

Example Input

Input 1:

 A = [1, 1, 3]
Input 2:

 A = [2, 4, 5]


Example Output

Output 1:

 1
Output 2:

 0


Example Explanation

Explanation 1:

 We can increase the value of 1st element by 1 in 1 step and will get all unique elements. i.e [2, 1, 3].
Explanation 2:

 All elements are distinct.

```

## Inversion Count in Array

```
Prompt: Inversion count in an array
Problem Description

Given an array of integers A. If i < j and A[i] > A[j], then the pair (i, j) is called an inversion of A. Find the total number of inversions of A modulo (109 + 7).

Problem Constraints
1 <= length of the array <= 105
1 <= A[i] <= 109
Input Format
The only argument given is the integer array A.
Output Format

Return the number of inversions of A modulo (109 + 7).

Example Input

Input 1:

A = [1, 3, 2]
Input 2:

A = [3, 4, 1, 2]


Example Output

Output 1:

1
Output 2:

4


Example Explanation

Explanation 1:

The pair (1, 2) is an inversion as 1 < 2 and A[1] > A[2]
Explanation 2:

The pair (0, 2) is an inversion as 0 < 2 and A[0] > A[2]
The pair (0, 3) is an inversion as 0 < 3 and A[0] > A[3]
The pair (1, 2) is an inversion as 1 < 2 and A[1] > A[2]
The pair (1, 3) is an inversion as 1 < 3 and A[1] > A[3]
```

``` go
package main

import (
	"fmt"
)

const MOD = 1000000007

func merge(arr []int, temp []int, left, mid, right int) int {
	i, j, k := left, mid+1, left
	invCount := 0

	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			invCount = (invCount + (mid - i + 1)) % MOD
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}

	for x := left; x <= right; x++ {
		arr[x] = temp[x]
	}

	return invCount % MOD
}

func mergeSort(arr []int, temp []int, left, right int) int {
	invCount := 0
	if left < right {
		mid := (left + right) / 2
		invCount = (invCount + mergeSort(arr, temp, left, mid)) % MOD
		invCount = (invCount + mergeSort(arr, temp, mid+1, right)) % MOD
		invCount = (invCount + merge(arr, temp, left, mid, right)) % MOD
	}
	return invCount % MOD
}

func inversionCount(arr []int) int {
	n := len(arr)
	temp := make([]int, n)
	return mergeSort(arr, temp, 0, n-1) % MOD
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(inversionCount(arr))
}


```

## Minimum Absolute Difference

```
Problem Description

Given an array of integers A, find and return the minimum value of | A [ i ] - A [ j ] | where i != j and |x| denotes the absolute value of x.



Problem Constraints

2 <= length of the array <= 100000

-109 <= A[i] <= 109



Input Format

The only argument given is the integer array A.



Output Format

Return the minimum value of | A[i] - A[j] |.



Example Input

Input 1:

 A = [1, 2, 3, 4, 5]
Input 2:

 A = [5, 17, 100, 11]


Example Output

Output 1:

 1
Output 2:

 6


Example Explanation

Explanation 1:

 The absolute difference between any two adjacent elements is 1, which is the minimum value.
Explanation 2:

 The minimum value is 6 (|11 - 5| or |17 - 11|).
```

``` go
import "sort"
import "math"
func solve(A []int )  (int) {
    sort.Ints(A)
    var ans int = math.MaxInt
    for i:=1;i<len(A);i++{
        ans=min(ans,A[i]-A[i-1])
    }
    return ans
}
```

## max chunks to make sorted

```
Problem Description

Given an array of integers A of size N that is a permutation of [0, 1, 2, ..., (N-1)], if we split the array into some number of "chunks" (partitions), and individually sort each chunk. After concatenating them in order of splitting, the result equals the sorted array.

What is the most number of chunks we could have made?






Problem Constraints

1 <= N <= 100000
0 <= A[i] < N



Input Format

The only argument given is the integer array A.



Output Format

Return the maximum number of chunks that we could have made.



Example Input

Input 1:

 A = [1, 2, 3, 4, 0]
Input 2:

 A = [2, 0, 1, 3]


Example Output

Output 1:

 1
Output 2:

 2
```

``` go

/**
 * @input A : Integer array
 * 
 * @Output Integer
 */

func solve(A []int )  (int) {
    i:=0
    ma:=0
    cnt:=0

    for _,e :=range A{
        ma=max(ma,e)
        if ma==i{
            cnt++
        }
        i++
    }
    return cnt

}
func max(A int, B int)int{
    //fmt.Print(A,B, " m ax ")
    if A>=B{
        return A
    }else{
        return B
    }
}

func invCount(A []int, l,r int)int{
    res:=0
    if r>l{
        mid:=(l+r)/2

        res+=invCount(A,l,mid)
        res+=invCount(A,mid+1,r)
        res+=merge(A,l,mid,r)
    }
    return res
}   

func merge(A []int ,l,y,r int )  (int) {
    
    temp:=make([]int,r-l+1)
    cnt:=0
    p1:=l
    p2:=y+1
    
    res:=0
    for p1<=y && p2<=r{
        if A[p1]<=A[p2]{
            temp[cnt]=A[p1]
            cnt++
            p1++
        }else{
            temp[cnt]=A[p2]
            cnt++
            p2++
            res+=y-p1
        }
    }

    for p1<=y{
        temp[cnt]=A[p1]
            cnt++
            p1++
    }
    for p2<=r{
        temp[cnt]=A[p2]
            cnt++
            p2++
    }
    for i:=0;i<=r-l+1;i++{
        A[i]=temp[i]
    }
    return res
}


```