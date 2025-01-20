# Bit Manipulation Truth Table

![alt text](logic_truth_table_1.png)

## Set ith Bit

Problem Description

You are given two integers A and B.
Set the A-th bit and B-th bit in 0, and return output in decimal Number System.

Note:
The bit positions are 0-indexed, which means that the least significant bit (LSB) has index 0.


Problem Constraints

0 <= A <= 30
0 <= B <= 30


Input Format

First argument A is an integer.
Second argument B is an integer.


Output Format

Return an integer.


Example Input

Input 1:
A = 3
B = 5
Input 2:
A = 4
B = 4


Example Output

Output 1:
40
Output 2:
16


Example Explanation

For Input 1:
The binary expression is 101000 which is 40 in decimal.
For Input 2:
The binary expression is 10000 which is 16 in decimal

``` go
func solve(A int , B int )  (int) {
    ans:= 0
    ans|= 1<<A
    
    ans = ans | 1<<B

    return ans
}

```

## Unset i-th bit

Problem Description

You are given two integers A and B.
If B-th bit in A is set, make it unset.
If B-th bit in A is unset, leave as it is.
Return the updated A value.

Note:
The bit position is 0-indexed, which means that the least significant bit (LSB) has index 0.


Problem Constraints

1 <= A <= 109
0 <= B <= 30


Input Format

First argument A is an integer.
Second argument B is an integer.


Output Format

Return an integer.


Example Input

Input 1:
A = 4
B = 1
Input 2:
A = 5
B = 2


Example Output

Output 1:
4
Output 2:
1


Example Explanation

For Input 1:
Given N = 4 which is 100 in binary. The 1-st bit is already unset
For Input 2:
Given N = 5 which is 101 in binary. We unset the 2-nd bit
It becomes 001 which is 1 in Decimal.

``` go
func solve(A int , B int )  (int) {
    ans:=0
    if A & (1<<B) > 0{
        ans = A &^(1<<B)
        return ans
    }else{
        return A
    }
}
```

## Check bit

Problem Description

You are given two integers A and B.
Return 1 if B-th bit in A is set
Return 0 if B-th bit in A is unset
Note:
The bit position is 0-indexed, which means that the least significant bit (LSB) has index 0.


Problem Constraints

1 <= A <= 109
0 <= B <= 30


Input Format

First argument A is an integer.
Second argument B is an integer.


Output Format

Return an integer.


Example Input

Input 1:
A = 4
B = 1
Input 2:
A = 5
B = 2


Example Output

Output 1:
0
Output 2:
1


Example Explanation

For Input 1:
Given N = 4 which is 100 in binary. The 1-st bit is unset
so we return 0
For Input 2:
Given N = 5 which is 101 in binary. The 2-nd bit is set
so we return 1

``` go
func solve(A int , B int )  (int) {
    if A & (1<<B) > 0{
        return 1
    }else{
        return 0
    }
}
```

## Number of 1 Bits

Problem Description

Write a function that takes an integer and returns the number of 1 bits present in its binary representation.


Problem Constraints

1 <= A <= 109


Input Format

First and only argument contains integer A


Output Format

Return an integer


Example Input

Input 1:
11
Input 2:
6


Example Output

Output 1:
3
Output 2:
2


Example Explanation

Explaination 1:
11 is represented as 1011 in binary.
Explaination 2:
6 is represented as 110 in binary.

``` go
func CountBits(A int) int {
	cnt := 0
	for A > 0 {
		n := A % 2
		if n == 1 {
			cnt++
		}
		A /= 2
	}
	return cnt
}

func NumSetBits(A int) int {
	cnt := 0

	for A > 0 {
		if A&1 != 0 {
			cnt++
		}
		A = A >> 1
	}
	return cnt
}
```

## Help From Sam

Problem Description

Alex and Sam are good friends. Alex is doing a lot of programming these days. He has set a target score of A for himself.
Initially, Alex's score was zero. Alex can double his score by doing a question, or Alex can seek help from Sam for doing questions that will contribute 1 to Alex's score. Alex wants his score to be precisely A. Also, he does not want to take much help from Sam.

Find and return the minimum number of times Alex needs to take help from Sam to achieve a score of A.


Problem Constraints

0 <= A <= 10^9


Input Format

The only argument given is an integer A.


Output Format

Return the minimum number of times help taken from Sam.


Example Input

Input 1:
A = 5
Input 2:

A = 3


Example Output

Output 1:
2
Output 2:

2


Example Explanation

Explanation 1:
Initial score : 0
Takes help from Sam, score : 1
Alex solves a question, score : 2
Alex solves a question, score : 4
Takes help from Sam, score: 5
Explanation 2:

Initial score : 0
Takes help from Sam, score : 1
Alex solves a question, score : 2
Takes help from Sam, score : 3

``` go
func solve(A int )  (int) {
     cnt:=0
    for i:=0;i<32;i++{
        if A & (1<<i)>0{
            cnt++
        }
    }
    return cnt
}

```

## Toggle i-th bit

Problem Description

You are given two integers A and B.
If B-th bit in A is set, make it unset
If B-th bit in A is unset, make it set
Return the updated A value


Problem Constraints

1 <= A <= 109
0 <= B <= 30


Input Format

First argument A is an integer.
Second argument B is an integer.


Output Format

Return an integer.


Example Input

Input 1:
A = 4
B = 1
Input 2:
A = 5
B = 2


Example Output

Output 1:
6
Output 2:
1


Example Explanation

For Input 1:
Given N = 4 which is 100 in binary. The 1-st bit is unset
so we make it set
For Input 2:
 
Given N = 5 which is 101 in binary. The 2-nd bit is set
so we make it unset

``` go
func solve(A int , B int )  (int) {
    return A ^ (1<<B)
}
```

## Unset x bits from right

Problem Description

Given an integer A. Unset B bits from the right of A in binary.

For example, if A = 93 and B = 4, the binary representation of A is 1011101.
If we unset the rightmost 4 bits, we get the binary number 1010000, which is equal to the decimal value 80.


Problem Constraints

1 <= A <= 1018
1 <= B <= 60


Input Format

The first argument is a single integer A.
The second argument is a single integer B.


Output Format

Return the number with B unset bits from the right.


Example Input
```
Input 1:-
A = 25
B = 3
Input 2:-
A = 37
B = 3


Example Output

Output 1:-
24
Output 2:-
32


Example Explanation

Explanation 1:-
A = 11001 to 11000
Explantio 2:-
A = 100101 to 100000
```

``` go
func solve(A int64 , B int )  (int64) {
    //var ans int64
    for i:=0;i<B;i++{
        if A & int64(1<<i)>0{
            A &^= int64(1<<i)
        }
    }
    return int64(A)
}
```

## Finding Good Days

Problem Description

Alex has a cat named Boomer. He decides to put his cat to the test for eternity.

He starts on day 1 with one stash of food unit, every next day, the stash doubles.

If Boomer is well behaved during a particular day, only then she receives food worth equal to the stash produced on that day.

Boomer receives a net worth of A units of food. What is the number of days she received the stash?



Problem Constraints

1 <= A <= 231-1



Input Format

First and only argument is an integer A.



Output Format

Return an integer denoting the number of days Boomer was well behaved.

```

Example Input

Input 1:

A = 5
Input 2:

A = 8


Example Output

Output 1:

 2
Output 2:

 1


Example Explanation

Explanation 1:

 To eat a total of 5 units of food, Boomer behaved normally on Day 1 and on the Day 3.
Explanation 2:

 To eat a total of 8 units of food, Boomer behaved normally only on day 4.
 ```

 ``` go
func solve(A int )  (int) {
    cnt:=0
    for i:=0;i<32;i++{
        if A & (1<<i)>0{
            cnt++
        }
    }
    return cnt
}
```
## Find nth Magic Number
Problem Description

Given an integer A, find and return the Ath magic number.

A magic number is defined as a number that can be expressed as a power of 5 or a sum of unique powers of 5.

First few magic numbers are 5, 25, 30(5 + 25), 125, 130(125 + 5), ….



Problem Constraints

1 <= A <= 5000



Input Format

The only argument given is integer A.



Output Format

Return the Ath magic number.

```

Example Input

Example Input 1:

 A = 3
Example Input 2:

 A = 10


Example Output

Example Output 1:

 30
Example Output 2:

 650


Example Explanation

Explanation 1:

 Magic Numbers in increasing order are [5, 25, 30, 125, 130, ...]
 3rd element in this is 30
Explanation 2:

 In the sequence shown in explanation 1, 10th element will be 650.
```
``` go
func solve(A int )  (int) {
    ans := 0
	pow := 5
	for A > 0 {
		r := A % 2
		//A = A / 2
		A = A >> 1
		ans += (r * pow)
		pow *= 5
	}
	return ans
}
```

not recomended as math lib usage
``` go
import "math"
func solve(A int )  (int) {
    var ans float64
   // power:=5
    for i:=0;i<32;i++{
       if A & (1<<i)>0{
           var j float64
        j=float64(i)
        ans+=math.Pow(5,(j+1))
        
        
       }
    
    }
    

    return int(ans)
}
```