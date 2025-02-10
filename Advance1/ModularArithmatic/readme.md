## Prime Modulo Inverse

```
Problem Description

Given two integers A and B. Find the value of A-1 mod B where B is a prime number and gcd(A, B) = 1.

A-1 mod B is also known as modular multiplicative inverse of A under modulo B.



Problem Constraints

1 <= A <= 109
1<= B <= 109
B is a prime number



Input Format

First argument is an integer A.
Second argument is an integer B.



Output Format

Return an integer denoting the modulor inverse



Example Input

Input 1:

 A = 3
 B = 5
Input 2:

 A = 6
 B = 23


Example Output

Output 1:

 2
Output 2:

 4


Example Explanation

Explanation 1:

 Let's say A-1 mod B = X, then (A * X) % B = 1.
 3 * 2 = 6, 6 % 5 = 1.
Explanation 2:

 Similarly, (6 * 4) % 23 = 1.


```

``` go
func power(x, y, p int) int{ 
    res := 1      // Initialize result 
    x = x % p  // Update x if it is more than or equal to p 
  
    for y > 0 { 
        // If y is odd, multiply x with result 
        if (y % 2 != 0) { 
            res = ((res % p) * (x % p)) % p
        } 
        y = y >> 1 
        x = (x * x) % p
    } 
    return res
}
func solve(A int, B int) (int) {
    return power (A, B - 2, B)
}

```

## Pair Sum divisible by M

```
Problem Description

Given an array of integers A and an integer B, find and return the number of pairs in A whose sum is divisible by B.

Since the answer may be large, return the answer modulo (109 + 7).

Note: Ensure to handle integer overflow when performing the calculations.


Problem Constraints

1 <= length of the array <= 100000
1 <= A[i] <= 109
1 <= B <= 106



Input Format

The first argument given is the integer array A.
The second argument given is the integer B.



Output Format

Return the total number of pairs for which the sum is divisible by B modulo (109 + 7).



Example Input

Input 1:

 A = [1, 2, 3, 4, 5]
 B = 2
Input 2:

 A = [5, 17, 100, 11]
 B = 28


Example Output

Output 1:

 4
Output 2:

 1


Example Explanation

Explanation 1:
 All pairs which are divisible by 2 are (1,3), (1,5), (2,4), (3,5). 
 So total 4 pairs.
Explanation 2:
 There is only one pair which is divisible by 28 is (17, 11)
 ```

 ``` go
package main

import (
	"fmt"
)

func countPairsModM(arr []int, M int) int {
	// Step 1: Compute frequency of remainders
	remainderCount := make(map[int]int)
	for _, num := range arr {
		remainder := num % M
		remainderCount[remainder]++
	}

	// Step 2: Count valid pairs
	count := 0

	// Case 1: Pairs with remainder 0 (A[i] + A[j] ≡ 0 (mod M))
	if remainderCount[0] > 1 {
		count += (remainderCount[0] * (remainderCount[0] - 1)) / 2
	}

	// Case 2: Pairs with remainder r and M-r
	for r := 1; r <= M/2; r++ {
		if r == M-r { // Special case when M is even
			count += (remainderCount[r] * (remainderCount[r] - 1)) / 2
		} else {
			count += remainderCount[r] * remainderCount[M-r]
		}
	}

	return count
}

func main() {
	arr := []int{2, 3, 7, 5, 8, 4} // Example array
	M := 5                           // Example M
	fmt.Println(countPairsModM(arr, M)) // Output: Count of valid pairs
}

 ```

##  Trailing Zeros in Factorial

 ```

Problem Description

Given an integer A, return the number of trailing zeroes in A! i.e., factorial of A.

Note: Your solution should run in logarithmic time complexity.


Problem Constraints

1 <= A <= 109


Input Format

First and only argument is a single integer A.



Output Format

Return a single integer denoting number of zeroes.



Example Input

Input 1

 A = 5
Input 2:

 A = 6


Example Output

Output 1:

 1
Output 2:

 1


Example Explanation

Explanation 1:

 A! = 120.
 Number of trailing zeros = 1. So, return 1.
Explanation 2:

 A! = 720 
 Number of trailing zeros = 1. So, return 1.
 ```

 ``` go
 func trailingZeroes(A int )  (int) {
    sum:=0
    for A/5>0{
        sum+=A/5
        A/=5
    }
    return sum
}
```

## Very large Power

```
Problem Description

Given two Integers A, B. You have to calculate (A ^ (B!)) % (1e9 + 7).

"^" means power,
"%" means mod, and
"!" means factorial.

Note: Ensure to handle integer overflow when performing the calculations.


Problem Constraints

1 <= A, B <= 5e5



Input Format

First argument is the integer A

Second argument is the integer B



Output Format

Return one integer, the answer to the problem



Example Input

Input 1:

A = 1
B = 1
Input 2:

A = 2
B = 2


Example Output

Output 1:

1
Output 2:

4


Example Explanation

Explanation 1:

 1! = 1. Hence 1^1 = 1.
Explanation 2:

 2! = 2. Hence 2^2 = 4.
```

## A,B and Modulo
```
Problem Description

Given two integers A and B, find the greatest possible positive integer M, such that A % M = B % M.



Problem Constraints

1 <= A, B <= 109
A != B



Input Format

The first argument is an integer A.
The second argument is an integer B.



Output Format

Return an integer denoting the greatest possible positive M.



Example Input

Input 1:

A = 1
B = 2
Input 2:

A = 5
B = 10


Example Output

Output 1:

1
Output 2:

5


Example Explanation

Explanation 1:

1 is the largest value of M such that A % M == B % M.
Explanation 2:

For M = 5, A % M = 0 and B % M = 0.

No value greater than M = 5, satisfies the condition.
```

``` go

func solve(A int , B int )  (int) {
    if A >= B{
        return A-B
    }
    return B-A

}

```

## Mod Sum

```
Problem Description

Given an array of integers A, calculate the sum of A [ i ] % A [ j ] for all possible i, j pairs. Return sum % (109 + 7) as an output.



Problem Constraints

1 <= length of the array A <= 105

1 <= A[i] <= 103
```

``` go
func solve(A []int )  (int) {
    n:=len(A)

    var mod int = 1e9+7

    var ct [1009]int

    for i:=0;i<n;i++{
        ct[A[i]]++
    }

    ans:=0

    for i:=1;i<=1000;i++{
        if ct[i]==0{
            continue
        }

        for j:=0;j<=1000;j++{
            if ct[j]==0{
                continue
            }

            val:=j%i
            temp:=ct[i]*ct[j]*val

            ans=((ans%mod)+(temp%mod))%mod
        }
        

    }
    return ans
}

```