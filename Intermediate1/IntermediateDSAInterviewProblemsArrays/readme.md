>Sum of even indices
Problem Description

You are given an array A of length N and Q queries given by the 2D array B of size Q*2. Each query consists of two integers B[i][0] and B[i][1].
For every query, the task is to calculate the sum of all even indices in the range A[B[i][0]…B[i][1]].

Note : Use 0-based indexing


Problem Constraints

1 <= N <= 105
1 <= Q <= 105
1 <= A[i] <= 100
0 <= B[i][0] <= B[i][1] < N


Input Format

First argument A is an array of integers.
Second argument B is a 2D array of integers.


Output Format

Return an array of integers.


Example Input

Input 1:
A = [1, 2, 3, 4, 5]
B = [   [0,2] 
        [1,4]   ]
Input 2:
A = [2, 1, 8, 3, 9]
B = [   [0,3] 
        [2,4]   ]


Example Output

Output 1:
[4, 8]
Output 2:
[10, 17]


Example Explanation

For Input 1:
The subarray for the first query is [1, 2, 3] whose sum of even indices is 4.
The subarray for the second query is [2, 3, 4, 5] whose sum of even indices is 8.
For Input 2:
The subarray for the first query is [2, 1, 8, 3] whose sum of even indices is 10.
The subarray for the second query is [8, 3, 9] whose sum of even indices is 17.

>Sum of odd indices

Problem Description

You are given an array A of length N and Q queries given by the 2D array B of size Q*2. Each query consists of two integers B[i][0] and B[i][1].
For every query, the task is to calculate the sum of all odd indices in the range A[B[i][0]…B[i][1]].

Note : Use 0-based indexing


Problem Constraints

1 <= N <= 105
1 <= Q <= 105
1 <= A[i] <= 100
0 <= B[i][0] <= B[i][1] < N


Input Format

First argument A is an array of integers.
Second argument B is a 2D array of integers.


Output Format

Return an array of integers.


Example Input

Input 1:
A = [1, 2, 3, 4, 5]
B = [   [0,2] 
        [1,4]   ]
Input 2:
A = [2, 1, 8, 3, 9]
B = [   [0,3] 
        [2,4]   ]


Example Output

Output 1:
[2, 6]
Output 2:
[4, 3]


Example Explanation

For Input 1:
The subarray for the first query is [1, 2, 3] whose sum of odd indices is 2.
The subarray for the second query is [2, 3, 4, 5] whose sum of odd indices is 6.
For Input 2:
The subarray for the first query is [2, 1, 8, 3] whose sum of odd indices is 4.
The subarray for the second query is [8, 3, 9] whose sum of odd indices is 3.

>Special Index

Problem Description

Given an array, arr[] of size N, the task is to find the count of array indices such that removing an element from these indices makes the sum of even-indexed and odd-indexed array elements equal.



Problem Constraints

1 <= N <= 105
-105 <= A[i] <= 105
Sum of all elements of A <= 109


Input Format

First argument contains an array A of integers of size N


Output Format

Return the count of array indices such that removing an element from these indices makes the sum of even-indexed and odd-indexed array elements equal.



Example Input

Input 1:
A = [2, 1, 6, 4]
Input 2:

A = [1, 1, 1]






Example Output

Output 1:
1
Output 2:

3






Example Explanation

Explanation 1:
Removing arr[1] from the array modifies arr[] to { 2, 6, 4 } such that, arr[0] + arr[2] = arr[1]. 
Therefore, the required output is 1. 
Explanation 2:

Removing arr[0] from the given array modifies arr[] to { 1, 1 } such that arr[0] = arr[1] 
Removing arr[1] from the given array modifies arr[] to { 1, 1 } such that arr[0] = arr[1] 
Removing arr[2] from the given array modifies arr[] to { 1, 1 } such that arr[0] = arr[1] 
Therefore, the required output is 3.

>Majority Element

Problem Description

Given an array of size N, find the majority element. The majority element is the element that appears more than floor(n/2) times.
You may assume that the array is non-empty and the majority element always exists in the array.



Problem Constraints

1 <= N <= 5*105
1 <= num[i] <= 109


Input Format

Only argument is an integer array.


Output Format

Return an integer.


Example Input

Input 1:
[2, 1, 2]
Input 2:
[1, 1, 1]


Example Output

Input 1:
2
Input 2:
1


Example Explanation

For Input 1:
2 occurs 2 times which is greater than 3/2.
For Input 2:
 1 is the only element in the array, so it is majority

 > N/3 Repeat Number <B>(Need to Solve)</B>
 Problem Description

You're given a read-only array of N integers. Find out if any integer occurs more than N/3 times in the array in linear time and constant additional space.
If so, return the integer. If not, return -1.

If there are multiple solutions, return any one.

Note: Read-only array means that the input array should not be modified in the process of solving the problem



Problem Constraints

1 <= N <= 7*105
1 <= A[i] <= 109


Input Format

The only argument is an integer array A.


Output Format

Return an integer.


Example Input

Input 1:
[1 2 3 1 1]
Input 2:
[1 2 3]


Example Output

Output 1:
1
Output 2:
-1


Example Explanation

Explanation 1:
1 occurs 3 times which is more than 5/3 times.
Explanation 2:
No element occurs more than 3 / 3 = 1 times in the array.

> Color of the last ball

You have 20 blue balls and 14 red balls in a bag. You put your hand in and remove 2 at a time.

If they’re of the same color, you add a blue ball to the bag.
If they’re of different colors, you add a red ball to the bag.
( Assume you have a big supply of blue & red balls for this purpose).

Note: When you take the two balls out, you don’t put them back in, so the number of balls in the bag keeps decreasing.

What will be the color of the last ball left in the bag?

>Complete Solution

There are three possible cases of removing the two balls.\
a) If we take off 1 RED and 1 BLUE, we actually will take off 1 BLUE.\
b)If we take off 2 RED, we actually will take off 2 RED (and add 1 BLUE).\
c) If we take off 2 BLUE, we actually will take off 1 BLUE.

So In the case of (a) or (c), we are only removing one blue ball, but we always take off red balls two by two.

If there are 14 (even) red balls, we can not have one red ball left in the bag so that the last ball will be blue.


>Color of the last ball II

You have 20 blue balls and 13 red balls in a bag. You put your hand in and remove two at a time.

If they’re of the same color, you add a blue ball to the bag.
If they’re of different colors, you add a red ball to the bag.
Assume you have an enormous supply of blue & red balls for this purpose.

Note: When you take the two balls out, you don’t put them back in, so the number of balls in the bag keeps decreasing.

What will be the color of the last ball left in the bag?

>Complete Solution

There are three possible cases of removing the two balls.\
a) If we take off 1 RED and 1 BLUE, we actually will take off 1 BLUE.\
b)If we take off 2 RED, we actually will take off 2 RED (and add 1 BLUE).\
c) If we take off 2 BLUE, we actually will take off 1 BLUE.

So In case of (a) or (c), we are only removing one blue ball, but we always take off red balls two by two.

Now as the no. of red balls is odd, there will be one single red ball in the bag with other blue balls, and whenever we remove 1 red and 1 blue ball, we end up taking off only the blue ball. So the red ball will be the last ball in the bag.