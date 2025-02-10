## Delete one

```
Problem Description

Given an integer array A of size N. You have to delete one element such that the GCD(Greatest common divisor) of the remaining array is maximum.

Find the maximum value of GCD.



Problem Constraints

2 <= N <= 105
1 <= A[i] <= 109



Input Format

First argument is an integer array A.



Output Format

Return an integer denoting the maximum value of GCD.



Example Input

Input 1:

 A = [12, 15, 18]
Input 2:

 A = [5, 15, 30]




Example Output

Output 1:

 6
Output 2:

 15




Example Explanation

Explanation 1:

 If you delete 12, gcd will be 3.
 If you delete 15, gcd will be 6.
 If you delete 18, gcd will 3.
 Maximum value of gcd is 6.
Explanation 2:

 If you delete 5, gcd will be 15.
 If you delete 15, gcd will be 5.
 If you delete 30, gcd will be 5.
 Maximum value of gcd is 15.
```

## Greatest Common Divisor

```
Problem Description

Given 2 non-negative integers A and B, find gcd(A, B)



GCD of 2 integers A and B is defined as the greatest integer 'g' such that 'g' is a divisor of both A and B. Both A and B fit in a 32 bit signed integer.

Note: DO NOT USE LIBRARY FUNCTIONS.





Problem Constraints

0 <= A, B <= 109



Input Format

First argument is an integer A.
Second argument is an integer B.



Output Format

Return an integer denoting the gcd(A, B).



Example Input

Input 1:

A = 4
B = 6
Input 2:

A = 6
B = 7


Example Output

Output 1:

 2
Output 2:

 1


Example Explanation

Explanation 1:

 2 divides both 4 and 6
Explanation 2:

 1 divides both 6 and 7
```

## PUBG

```
Problem Description

There are N players, each with strength A[i]. when player i attack player j, player j strength reduces to max(0, A[j]-A[i]). When a player's strength reaches zero, it loses the game, and the game continues in the same manner among other players until only 1 survivor remains.

Can you tell the minimum health last surviving person can have?



Problem Constraints

1 <= N <= 100000

1 <= A[i] <= 1000000



Input Format

First and only argument of input contains a single integer array A.



Output Format

Return a single integer denoting minimum health of last person.



Example Input

Input 1:

 A = [6, 4]
Input 2:

 A = [2, 3, 4]


Example Output

Output 1:

 2
Output 2:

 1


Example Explanation

Explanation 1:

 Given strength array A = [6, 4]
 Second player attack first player, A =  [2, 4]
 First player attack second player twice. [2, 0]
Explanation 2:

 Given strength array A = [2, 3, 4]
 First player attack third player twice. [2, 3, 0]
 First player attack second player. [2, 1, 0]
 Second player attack first player twice. [0, 1, 0]
```
``` go

func solve(A []int )  (int) {
        ans:=A[0]
    for i:=0;i<len(A);i++{
        ans=gcd(ans,A[i])
    }

    return ans
}

func gcd(A,B int)int{
    if A==0{
        return B
    }
    return gcd(B%A,A)
}

```

## GCD CMPL

```
In the following C++ function, let n >= m.

    int gcd(int n, int m) {
            if (n%m ==0) return m;
            if (n < m) swap(n, m);
            while (m > 0) {
                n = n%m;
                swap(n, m);
            }
            return n;
    }
What is the time complexity of the above function assuming n > m?.
Θ symbol represents theta notation and Ω symbol represents omega notation.
```
### Θ(logn)


## Largest Coprime Divisor

```
Problem Description

You are given two positive numbers A and B . You need to find the maximum valued integer X such that:

X divides A i.e. A % X = 0
X and B are co-prime i.e. gcd(X, B) = 1


Problem Constraints

1 <= A, B <= 109



Input Format

First argument is an integer A.
Second argument is an integer B.



Output Format

Return an integer maximum value of X as descibed above.



Example Input

Input 1:

 A = 30
 B = 12
Input 2:

 A = 5
 B = 10


Example Output

Output 1:

 5
Output 2:

 1


Example Explanation

Explanation 1:

 All divisors of A are (1, 2, 3, 5, 6, 10, 15, 30). 
 The maximum value is 5 such that A%5 == 0 and gcd(5,12) = 1
Explanation 2:

 1 is the only value such that A%5 == 0 and gcd(1,10) = 1

```

