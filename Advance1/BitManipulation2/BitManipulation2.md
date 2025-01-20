## Single Number

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


```
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
```
``` go
func singleNumber(A []int )  (int) {
    ans:=0
    for i:=0;i<len(A);i++{
        ans^=A[i]
    }
    return ans
}
```

## Single Number II

Problem Description

Given an array of integers, every element appears thrice except for one, which occurs once.

Find that element that does not appear thrice.

NOTE: Your algorithm should have a linear runtime complexity.

Could you implement it without using extra memory?

```
Problem Constraints

2 <= A <= 5*106

0 <= A <= INTMAX



Input Format

First and only argument of input contains an integer array A.



Output Format

Return a single integer.



Example Input

Input 1:

 A = [1, 2, 4, 3, 3, 2, 2, 3, 1, 1]
Input 2:

 A = [0, 0, 0, 1]


Example Output

Output 1:

 4
Output 2:

 1


Example Explanation

Explanation 1:

 4 occurs exactly once in Input 1.
 1 occurs exactly once in Input 2.
 ```

 ## Single Number III


Problem Description

Given an array of positive integers A, two integers appear only once, and all the other integers appear twice.
Find the two integers that appear only once.

Note: Return the two numbers in ascending order.


Problem Constraints

2 <= |A| <= 100000
1 <= A[i] <= 109



Input Format

The first argument is an array of integers of size N.



Output Format

Return an array of two integers that appear only once.

```
Example Input

Input 1:
A = [1, 2, 3, 1, 2, 4]
Input 2:

A = [1, 2]


Example Output

Output 1:
[3, 4]
Output 2:

[1, 2]


Example Explanation

Explanation 1:
3 and 4 appear only once.
Explanation 2:

1 and 2 appear only once.

```
## Find Two Missing Numbers

Problem Description

Given an array A of length N where all the elements are distinct and are in the range [1, N+2].

Two numbers from the range [1, N+2] are missing from the array A. Find the two missing numbers.



Problem Constraints

1 <= N <= 105

1 <= A[i] <= N+2

The elements of array A are distinct

```
Input Format

Only argument A is an array of integers



Output Format

Return a sorted array of size 2 denoting the missing elements.



Example Input

Input 1:
A = [3, 2, 4]
Input 2:
A = [5, 1, 3, 6]


Example Output

Output 1:
[1, 5]
Output 2:
[2, 4]


Example Explanation

For Input 1:
The missing numbers are 1 and 5.
For Input 2:
The missing numbers are 2 and 4.
```

## Maximum AND Pair

Problem Description

Given an array A. For every pair of indices i and j (i != j), find the maximum A[i] & A[j].


Problem Constraints

1 <= len(A) <= 105
1 <= A[i] <= 109


Input Format

The first argument is an integer array A.


Output Format

Return a single integer that is the maximum A[i] & A[j].

```
Example Input

Input 1:-
A = [53, 39, 88]
Input 2:-
A = [38, 44, 84, 12] 


Example Output

Output 1:-
37
Output 2:-
36


Example Explanation

Explanation 1:-
53 & 39 = 37
39 & 88 = 0
53 & 88 = 16
Maximum among all these pairs is 37
Explanation 2:-
Maximum bitwise and among all pairs is (38, 44) = 36
```

## Strange Equality

Problem Description

Given an integer A.
Two numbers, X and Y, are defined as follows:

X is the greatest number smaller than A such that the XOR sum of X and A is the same as the sum of X and A.
Y is the smallest number greater than A, such that the XOR sum of Y and A is the same as the sum of Y and A.
Find and return the XOR of X and Y.

NOTE 1: XOR of X and Y is defined as X ^ Y where '^' is the BITWISE XOR operator.

NOTE 2: Your code will be run against a maximum of 100000 Test Cases.



Problem Constraints

1 <= A <= 109



Input Format

First and only argument is an integer A.



Output Format

Return an integer denoting the XOR of X and Y.



Example Input

A = 5


Example Output

10


Example Explanation

The value of X will be 2 and the value of Y will be 8. The XOR of 2 and 8 is 10.


``` go
func solve(A int )  (int) {
    bit:=0
    x:=0
    y:=0

    for A !=0{
        if A%2==0{
        x|=(1<<bit)
        }
        A/=2
        bit++
    }
    y=(1<<bit)

    return x^y
}
```

## SUBARRAY OR
Problem Description

You are given an array of integers A of size N.

The value of a subarray is defined as BITWISE OR of all elements in it.

Return the sum of value of all subarrays of A % 109 + 7.



Problem Constraints

1 <= N <= 105

1 <= A[i] <= 108



Input Format

The first argument given is the integer array A.



Output Format

Return the sum of Value of all subarrays of A % 109 + 7.


```
Example Input

Input 1:

 A = [1, 2, 3, 4, 5]
Input 2:

 A = [7, 8, 9, 10]


Example Output

Output 1:

 71
Output 2:

 110


Example Explanation

Explanation 1:

 Value([1]) = 1
 Value([1, 2]) = 3
 Value([1, 2, 3]) = 3
 Value([1, 2, 3, 4]) = 7
 Value([1, 2, 3, 4, 5]) = 7
 Value([2]) = 2
 Value([2, 3]) = 3
 Value([2, 3, 4]) = 7
 Value([2, 3, 4, 5]) = 7
 Value([3]) = 3
 Value([3, 4]) = 7
 Value([3, 4, 5]) = 7
 Value([4]) = 4
 Value([4, 5]) = 5
 Value([5]) = 5
 Sum of all these values = 71
Explanation 2:

 Sum of value of all subarray is 110.
```

## Min XOR value
Problem Description

Given an integer array A of N integers, find the pair of integers in the array which have minimum XOR value. Report the minimum XOR value.


```
Problem Constraints

2 <= length of the array <= 100000
0 <= A[i] <= 109



Input Format

First and only argument of input contains an integer array A.



Output Format

Return a single integer denoting minimum xor value.



Example Input

Input 1:

 A = [0, 2, 5, 7]
Input 2:

 A = [0, 4, 7, 9]


Example Output

Output 1:

 2
Output 2:

 3


Example Explanation

Explanation 1:

 0 xor 2 = 2
```

## Sum of Xor of all Pairs

```
Problem Description

Given an array A of N integers. Find the sum of bitwise XOR all pairs of numbers in the array.

Since the answer can be large, return the remainder after the dividing the answer by 109+7.



Problem Constraints

1 <= N <= 105

1 <= A[i] < 109



Input Format

Only argument A is an array of integers



Output Format

Return an integer denoting the sum of xor of all pairs of number in the array.



Example Input

Input 1:
A = [1, 2, 3]
Input 2:
A = [3, 4, 2]


Example Output

Output 1:
6
Output 2:
14


Example Explanation

For Input 1:
Pair    Xor
{1, 2}  3
{1, 3}  2
{2, 3}  1
Sum of xor of all pairs = 3 + 2 + 1 = 6.
For Input 2:
Pair    Xor
{3, 4}  7
{3, 2}  1
{4, 2}  6
Sum of xor of all pairs = 7 + 1 + 6 = 14.

```

``` go
func solve(A []int) int {
    const MOD int = 1e9 + 7
    xorSum := 0

    for i := 0; i < len(A); i++ {
        for j := i + 1; j < len(A); j++ {
            xorSum += A[i] ^ A[j]
            xorSum %= MOD
        }
    }

    return xorSum
}
```

``` go
func solve(A []int) int {
    const MOD = 1000000007
    n := len(A)
    result := 0

    // Iterate over each bit position (0 to 31)
    for b := 0; b < 32; b++ {
        count1 := 0

        // Count numbers with the b-th bit set
        for _, num := range A {
            if (num & (1 << b)) != 0 {
                count1++
            }
        }

        count0 := n - count1 // Numbers with b-th bit unset
        contribution := (count1 * count0 % MOD) * (1 << b) % MOD
        result = (result + contribution) % MOD
    }

    return result
}

```