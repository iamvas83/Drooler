> Practice Questions

Given two binary numbers A = 1001011 and B = 11001001. What is their sum?


Binary to Decimal - III

Which of the following is the correct conversion of 10101101(base 2) to decimal?


Decimal to Binary
Which of the following is a correct conversion of 76 (base 10) to binary?

Bit Operations
Given a = 1010011 and b = 1001001, their OR, XOR and AND are -

If a&1 = 1, then a is?

What are the values of a&a, a|a, a^a?
sol- a&a = a , a|a = a , a^a = 0

Any base to decimal
Problem Description

You are given a number A. You are also given a base B. A is a number on base B.
You are required to convert the number A into its corresponding value in decimal number system.


Problem Constraints

0 <= A <= 109
2 <= B <= 9


Input Format

First argument A is an integer.
Second argument B is an integer.

Output Format

Return an integer.


Example Input

Input 1:
A = 1010
B = 2
Input 2:
A = 22 
B = 3

Example Output

Output 1:
10
Output 2:
8


Example Explanation

For Input 1:
The decimal 10 in base 2 is 1010.
For Input 2:
The decimal 8 in base 3 is 22.

>Decimal to Any Base

Problem Description

Given a decimal number A and a base B, convert it into its equivalent number in base B.


Problem Constraints

0 <= A <= 512
2 <= B <= 10


Input Format

The first argument will be decimal number A.
The second argument will be base B.


Output Format

Return the conversion of A in base B.


Example Input

Input 1:
A = 4
B = 3
Input 2:
A = 4
B = 2 


Example Output

Output 1:
11
Output 2:
100


Example Explanation

Explanation 1:
Decimal number 4 in base 3 is 11.
Explanation 2:
Decimal number 4 in base 2 is 100. 

>Single Number

Problem Description

Given an array of integers A, every element appears twice except for one. Find that integer that occurs once.

NOTE: Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?



Problem Constraints

1 <= |A| <= 2000000

0 <= A[i] <= INTMAX



Input Format

The first and only argument of input contains an integer array A.



Output Format

Return a single integer denoting the single element.



Example Input

Input 1:

 A = [1, 2, 2, 3, 1]
Input 2:

 A = [1, 2, 2]


Example Output

Output 1:

 3
Output 2:

 1


Example Explanation

Explanation 1:

3 occurs once.
Explanation 2:

1 occurs once.

>Subarrays with Bitwise OR 1

Problem Description

Given an array B of length A with elements 1 or 0. Find the number of subarrays such that the bitwise OR of all the elements present in the subarray is 1.
Note : The answer can be large. So, return type must be long.



Problem Constraints

1 <= A <= 105


Input Format

The first argument is a single integer A.
The second argument is an integer array B.

Output Format

Return the number of subarrays with bitwise array 1.


Example Input

Input 1:
 A = 3
B = [1, 0, 1]
Input 2:
 A = 2
B = [1, 0]


Example Output

Output 1:
5
Output2:
2

Example Explanation

Explanation 1:
The subarrays are :- [1], [0], [1], [1, 0], [0, 1], [1, 0, 1]
Except the subarray [0] all the other subarrays has a Bitwise OR = 1
Explanation 2:
The subarrays are :- [1], [0], [1, 0]
Except the subarray [0] all the other subarrays has a Bitwise OR = 1


What happens to the 0-th bit in a when we perform a = a^1 ?

ans-it gets toggled