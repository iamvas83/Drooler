## Count Subarray Zero Sum

Problem Description

Given an array A of N integers.

Find the count of the subarrays in the array which sums to zero. Since the answer can be very large, return the remainder on dividing the result with 109+7


Problem Constraints

1 <= N <= 105
-109 <= A[i] <= 109


Input Format

Single argument which is an integer array A.


Output Format

Return an integer.


Example Input

Input 1:

 A = [1, -1, -2, 2]
Input 2:

 A = [-1, 2, -1]


Example Output

Output 1:
3

Output 2:
1

Example Explanation

Explanation 1:

 The subarrays with zero sum are [1, -1], [-2, 2] and [1, -1, -2, 2].

Explanation 2:

 The subarray with zero sum is [-1, 2, -1].

``` go
    func solve(A []int )  (int) {
    m := make(map[int]int)
	sum, count := 0, 0
	m[0] = 1
	for i := 0; i < len(A); i++ {
		sum += A[i]
		if _, ok := m[sum]; ok {
			count += m[sum]
			m[sum] = m[sum] + 1
		} else {
			m[sum] = 1
		}
	}

	return count % 1000000007
}

```

## Common Elements

Problem Description

Given two integer arrays, A and B of size N and M, respectively. Your task is to find all the common elements in both the array.


NOTE:


Each element in the result should appear as many times as it appears in both arrays.
The result can be in any order.


Problem Constraints

1 <= N, M <= 105

1 <= A[i] <= 109



Input Format

First argument is an integer array A of size N.

Second argument is an integer array B of size M.



Output Format

Return an integer array denoting the common elements.



Example Input

Input 1:

 A = [1, 2, 2, 1]
 B = [2, 3, 1, 2]
Input 2:

 A = [2, 1, 4, 10]
 B = [3, 6, 2, 10, 10]


Example Output

Output 1:

 [1, 2, 2]
Output 2:

 [2, 10]


Example Explanation

Explanation 1:

 Elements (1, 2, 2) appears in both the array. Note 2 appears twice in both the array.
Explantion 2:

 Elements (2, 10) appears in both the array.

## Distinct Numbers in Window

Problem Description

You are given an array of N integers, A1, A2 ,..., AN and an integer B. Return the of count of distinct numbers in all windows of size B.

Formally, return an array of size N-B+1 where i'th element in this array contains number of distinct elements in sequence Ai, Ai+1 ,..., Ai+B-1.

NOTE: if B > N, return an empty array.

Problem Constraints

1 <= N <= 106

1 <= A[i] <= 109


Input Format

First argument is an integer array A
Second argument is an integer B.



Output Format

Return an integer array.



Example Input

Input 1:

 A = [1, 2, 1, 3, 4, 3]
 B = 3
Input 2:

 A = [1, 1, 2, 2]
 B = 1


Example Output

Output 1:

 [2, 3, 3, 2]
Output 2:

 [1, 1, 1, 1]


Example Explanation

Explanation 1:

 A=[1, 2, 1, 3, 4, 3] and B = 3
 All windows of size B are
 [1, 2, 1]
 [2, 1, 3]
 [1, 3, 4]
 [3, 4, 3]
 So, we return an array [2, 3, 3, 2].
Explanation 2:

 Window size is 1, so the output array is [1, 1, 1, 1].

 ##  Subarray Sum Equals K

 Problem Description

Given an array of integers A and an integer B.
Find the total number of subarrays having sum equals to B.


Problem Constraints

 1 <= length of the array <= 50000
-1000 <= A[i] <= 1000


Input Format

The first argument given is the integer array A.
The second argument given is integer B.


Output Format

Return the total number of subarrays having sum equals to B.


Example Input

Input 1:
A = [1, 0, 1]
B = 1
Input 2:
A = [0, 0, 0]
B = 0


Example Output

Output 1:
4
Output 2:
6


Example Explanation

Explanation 1:
[1], [1, 0], [0, 1] and [1] are four subarrays having sum 1.
Explanation 1:
All the possible subarrays having sum 0.


## Pair With Given Difference

Problem Description

Given an one-dimensional unsorted array A containing N integers.

You are also given an integer B, find if there exists a pair of elements in the array whose difference is B.

Return 1 if any such pair exists else return 0.



Problem Constraints

1 <= N <= 105
-103 <= A[i] <= 103
-105 <= B <= 105


Input Format

First argument is an integer array A of size N.

Second argument is an integer B.



Output Format

Return 1 if any such pair exists else return 0.



Example Input

Input 1:

 A = [5, 10, 3, 2, 50, 80]
 B = 78
Input 2:

 A = [-10, 20]
 B = 30
Input 2=3:

 A = [-100, 20, 80]
 B = 9


Example Output

Output 1:

 1
Output 2:

 1
Output 3:

 0


Example Explanation

Explanation 1:

 Pair (80, 2) gives a difference of 78.
Explanation 2:

 Pair (20, -10) gives a difference of 30 i.e 20 - (-10) => 20 + 10 => 30
Explanation 3:

 There is no such pairs, that has a difference of 9.

 ``` go

 func solve(A []int , B int )  (int) {
    m:=map[int]struct{}{}
    
    for i:=0;i<len(A);i++{
        if _,ok:=m[B+A[i]];ok{
            return 1
        }else if _,ok:=m[A[i]-B];ok{
            return 1
           // m[A[i]]=struct{}{}
        }else{
            m[A[i]] = struct{}{}
        }
    }
    return 0
}
```

## Count Pair Sum

Problem Description

You are given an array A of N integers and an integer B. Count the number of pairs (i,j) such that A[i] + A[j] = B and i ≠ j.

Since the answer can be very large, return the remainder after dividing the count with 109+7.

Note - The pair (i,j) is same as the pair (j,i) and we need to count it only once.


Problem Constraints

1 <= N <= 105
1 <= A[i] <= 109
1 <= B <= 109


Input Format

First argument A is an array of integers and second argument B is an integer.


Output Format

Return an integer.


Example Input

Input 1:

A = [3, 5, 1, 2]
B = 8
Input 2:

A = [1, 2, 1, 2]
B = 3


Example Output

Output 1:

1
Output 2:

4


Example Explanation

Example 1:

The only pair is (1, 2) which gives sum 8
Example 2:

The pair which gives sum as 3 are (1, 2), (1, 4), (2, 3) and (3, 4).

``` go

func solve(A []int , B int )  (int) {
    m:=map[int]int{}

    res:=0
    for i:=0;i<len(A);i++{
        if _,ok:=m[B-A[i]];ok{
            res=(res+m[B-A[i]])%100000007  
            
        }else{
            if _, ok := m[A[i]]; ok {
			m[A[i]] = m[A[i]] + 1
		} else {
			m[A[i]] = 1
		}
        }
    }   
    return res
}

```

## Check Pair Sum

Problem Description

Given an Array of integers B, and a target sum A.
Check if there exists a pair (i,j) such that Bi + Bj = A and i!=j.


Problem Constraints

1 <= Length of array B <= 105
0 <= Bi <= 109
0 <= A <= 109


Input Format

First argument A is the Target sum, and second argument is the array B


Output Format

Return an integer value 1 if there exists such pair, else return 0.


Example Input

Input 1:








A = 8   
B = [3, 5, 1, 2, 1, 2]
Input 2:

A = 21   
B = [9, 10, 7, 10, 9, 1, 5, 1, 5]









Example Output

Output 1:








1
Output 2:

0









Example Explanation

Example 1:
It is possible to obtain sum 8 using 3 and 5.
Example 2:
There is no such pair exists.

``` go

func solve(A int , B []int )  (int) {
    m:=map[int]struct{}{}

    for i:=0;i<len(B);i++{
        if _,ok:=m[A-B[i]];ok{
            return 1
        }else{
            m[B[i]]=struct{}{}
        }
    }
    return 0
}

```


## Pairs With Given Xor

Problem Description

Given an integer array A containing N distinct integers.




Find the number of unique pairs of integers in the array whose XOR is equal to B.

NOTE:

Pair (a, b) and (b, a) is considered to be the same and should be counted once.





Problem Constraints

1 <= N <= 105

1 <= A[i], B <= 107



Input Format

The first argument is an integer array A.







The second argument is an integer B.









Output Format

Return a single integer denoting the number of unique pairs of integers in the array A whose XOR is equal to B.



Example Input

Input 1:

 A = [5, 4, 10, 15, 7, 6]
 B = 5
Input 2:

 A = [3, 6, 8, 10, 15, 50]
 B = 5


Example Output

Output 1:

 1
Output 2:

 2


Example Explanation

Explanation 1:

 (10 ^ 15) = 5
Explanation 2:

 (3 ^ 6) = 5 and (10 ^ 15) = 5 

 ``` go

 func solve(A []int , B int )  (int) {
    m:=map[int]int{}
    res:=0
    for i:=0;i<len(A);i++{
        if _,ok:=m[B^A[i]];ok{
          res+=m[B^A[i]]  
        }
        m[A[i]]=m[A[i]]+1
    }
    return res
}

```

## Is Dictionary?

Problem Description

Surprisingly, in an alien language, they also use English lowercase letters, but possibly in a different order. The order of the alphabet is some permutation of lowercase letters.




Given an array of words A of size N written in the alien language, and the order of the alphabet denoted by string B of size 26, return 1 if and only if the given words are sorted lexicographically in this alien language else, return 0.






Problem Constraints

1 <= N, length of each word <= 105

Sum of the length of all words <= 2 * 106






Input Format

The first argument is a string array A of size N.




The second argument is a string B of size 26, denoting the order.






Output Format

Return 1 if and only if the given words are sorted lexicographically in this alien language else, return 0.



Example Input

Input 1:

 A = ["hello", "scaler", "interviewbit"]
 B = "adhbcfegskjlponmirqtxwuvzy"
Input 2:

 A = ["fine", "none", "bush"]
 B = "qwertyuiopasdfghjklzxcvbnm"


Example Output

Output 1:

 1
Output 2:

 0


Example Explanation

Explanation 1:

 The order shown in string B is: h < s < i (adhbcfegskjlponmirqtxwuvzy) for the given words. 
 So, Return 1.
Explanation 2:

 "none" should be present after "bush", Since b < n (qwertyuiopasdfghjklzxcvbnm). 
 So, Return 0.



