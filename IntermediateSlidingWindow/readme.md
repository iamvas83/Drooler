> <B>Subarray with given sum and length</B>

Problem Description

Given an array A of length N. Also given are integers B and C.

Return 1 if there exists a subarray with length B having sum C and 0 otherwise

Problem Constraints

1 <= N <= 10<sup>5</sup>

1 <= A[i] <= 10<sup>4</sup>

1 <= B <= N

1 <= C <= 10<sup>9</sup>

Input Format

First argument A is an array of integers.

The remaining arguments B and C are integers

Output Format

Return 1 if such a subarray exist and 0 otherwise

Example Input

Input 1:


A = [4, 3, 2, 6, 1]
B = 3
C = 11
Input 2:

A = [4, 2, 2, 5, 1]
B = 4
C = 6

Example Output

Output 1:
1
Output 2:

0

Example Explanation

Explanation 1:


The subarray [3, 2, 6] is of length 3 and sum 11.


Explanation 2:


There are no such subarray.


>$$!Minimum\ Swaps

Problem Description

Given an array of integers A and an integer B, find and return the minimum number of swaps required to bring all the numbers less than or equal to B together.

Note: It is possible to swap any two elements, not necessarily consecutive.

Problem Constraints

1 <= length of the array <= 100000
-10<sup>9</sup>  <= A[i], B <= 10<sup>9</sup>





><B>Spiral Order Matrix II</B>

Problem Description

Given an integer A, generate a square matrix filled with elements from 1 to A2 in spiral order and return the generated square matrix.

Problem Constraints

1 <= A <= 1000

Input Format

First and only argument is integer A


Output Format

Return a 2-D matrix which consists of the elements added in spiral order.

Example Input

Input 1:

1
Input 2:

2
Input 3:

5

Example Output

Output 1:

[ [1] ]
Output 2:

[ [1, 2],<br> 
  [4, 3] ]

Output 3:

[ [1,   2,  3,  4, 5],<br> 
  [16, 17, 18, 19, 6],<br>
  [15, 24, 25, 20, 7],<br>
  [14, 23, 22, 21, 8],<br> 
  [13, 12, 11, 10, 9] ]
