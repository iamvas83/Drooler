>Mod Array

Problem Description

You are given a large number in the form of a array A of size N where each element denotes a digit of the number.
You are also given a number B. You have to find out the value of A % B and return it.



Problem Constraints

1 <= N <= 105
0 <= Ai <= 9
1 <= B <= 109


Input Format

The first argument is an integer array A.
The second argument is an integer B.


Output Format

Return a single integer denoting the value of A % B.


Example Input

Input 1:
A = [1, 4, 3]
B = 2
Input 2:

A = [4, 3, 5, 3, 5, 3, 2, 1]
B = 47


Example Output

Output 1:
1
Output 2:

20


Example Explanation

Explanation 1:
143 is an odd number so 143 % 2 = 1.
Explanation 2:

43535321 % 47 = 20

>Divisibility by 3

Problem Description

Given a number in the form of an array A of size N. Each of the digits of the number is represented by A[i]. Check if the number is divisible by 3.


Problem Constraints

1 <= N <= 105

0 <= A[i] <= 9

A[1] ≠ 0



Input Format

Given an integer array representing the number



Output Format

Return 1 if the number is divisible by 3 and return 0 otherwise.



Example Input

Input 1:
A = [1, 2, 3]
Input 2:
A = [1, 0, 0, 1, 2]


Example Output

Output 1:
1
Output 2:
0


Example Explanation

For Input 1:
The number 123 is divisible by 3.
For Input 2:
The number 10012 is not divisible by 3.

## Brute Force ##

```go
    func ModArray(A []int, B int) int {
	ans := 0
	power := 1

	n := len(A)

	for i := n - 1; i > 0; i-- {
		ans = (ans + (A[i]%B*power%B)%B) % B
		power = (power * 10) % B
	}
	return ans
}
```