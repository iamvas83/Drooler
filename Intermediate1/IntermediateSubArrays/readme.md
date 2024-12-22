Total number of subarrays = (i+1)*(n-i)
i-> index
n -> number of elements in Array

possible start points =0,1,2,..i: i+1
possible end points = i,i+1,i+2,...n-1: n-i

<B>Total number of subarrays = (i+1)*(n-i)</B>

<B>Contribution of A[i] = A[i]\*(i+1)*(n-i)</B>

> MaxSumSubArrayEasy

Problem Description

You are given an integer array C of size A. Now you need to find a subarray (contiguous elements) so that the sum of contiguous elements is maximum.
But the sum must not exceed B.


Problem Constraints

1 <= A <= 103
1 <= B <= 109
1 <= C[i] <= 106

Input Format

The first argument is the integer A.
The second argument is the integer B.
The third argument is the integer array C.

Output Format

Return a single integer which denotes the maximum sum.

Example Input

Input 1:
A = 5
B = 12
C = [2, 1, 3, 4, 5]
Input 2:

A = 3
B = 1
C = [2, 2, 2]

Example Output
Output 1:
12
Output 2:
0

Example Explanation

Explanation 1:
We can select {3,4,5} which sums up to 12 which is the maximum possible sum.
Explanation 2:

All elements are greater than B, which means we cannot select any subarray.
Hence, the answer is 0.


> Q. Sum of All Subarrays

<B>Problem Description</B>

You are given an integer array A of length N.
You have to find the sum of all subarray sums of A.
More formally, a subarray is defined as a contiguous part of an array which we can obtain by deleting zero or more elements from either end of the array.
A subarray sum denotes the sum of all the elements of that subarray.

Note : Be careful of integer overflow issues while calculations. Use appropriate datatypes

Problem Constraints

1 <= N <= 105
1 <= Ai <= 10 4

Input Format

The first argument is the integer array A.

Output Format

Return a single integer denoting the sum of all subarray sums of the given array.

Example Input

Input 1:
A = [1, 2, 3]
Input 2:

A = [2, 1, 3]

Example Output

Output 1:
20
Output 2:
19

Example Explanation

Explanation 1:
The different subarrays for the given array are: [1], [2], [3], [1, 2], [2, 3], [1, 2, 3].
Their sums are: 1 + 2 + 3 + 3 + 5 + 6 = 20
Explanation 2:

The different subarrays for the given array are: [2], [1], [3], [2, 1], [1, 3], [2, 1, 3].
Their sums are: 2 + 1 + 3 + 3 + 4 + 6 = 19