## Find Fibonacci - II

Problem Description

The Fibonacci numbers are the numbers in the following integer sequence.



0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ……..

In mathematical terms, the sequence Fn of Fibonacci numbers is defined by the recurrence relation:

Fn = Fn-1 + Fn-2

Given a number A, find and return the Ath Fibonacci Number using recursion.

Given that F0 = 0 and F1 = 1.





Problem Constraints

0 <= A <= 20



Input Format

First and only argument is an integer A.



Output Format

Return an integer denoting the Ath term of the sequence.



Example Input

Input 1:

 A = 2
Input 2:

 A = 9


Example Output

Output 1:

 1
Output 2:

 34


Example Explanation

Explanation 1:

 f(2) = f(1) + f(0) = 1
Explanation 2:

 f(9) = f(8) + f(7) = 21 + 13  = 34

 ## Find Factorial!


Problem Description

Write a program to find the factorial of the given number A using recursion.

Note: The factorial of a number N is defined as the product of the numbers from 1 to N.


Problem Constraints

0 <= A <= 12



Input Format

First and only argument is an integer A.



Output Format

Return an integer denoting the factorial of the number A.



Example Input

Input 1:

 A = 4
Input 2:

 A = 1


Example Output

Output 1:

 24
Output 2:

 1


Example Explanation

Explanation 1:

 Factorial of 4 = 4 * 3 * 2 * 1 = 24
Explanation 2:

 Factorial of 1 = 1

 ``` go

 func solve(A int )  (int) {
    if A==1{
        return 1
    }
    return A*solve(A-1)
}

```

## Check Palindrome using Recursion

Problem Description

Write a recursive function that checks whether string A is a palindrome or Not.
Return 1 if the string A is a palindrome, else return 0.

Note: A palindrome is a string that's the same when read forward and backward.






Problem Constraints

1 <= |A| <= 50000




String A consists only of lowercase letters.






Input Format

The first and only argument is a string A.



Output Format

Return 1 if the string A is a palindrome, else return 0.



Example Input

Input 1:

 A = "naman"
Input 2:

 A = "strings"


Example Output

Output 1:

 1
Output 2:

 0


Example Explanation

Explanation 1:

 "naman" is a palindomic string, so return 1.
Explanation 2:

 "strings" is not a palindrome, so return 0.

 ``` go

 func solve(A string )  (int) {
    x:=0
    y:=len(A)-1

    if isPal(A,x,y){
        return 1
    }else{
        return 0
    }

}

func isPal(A string, x int, y int)bool{
    if x>y{
        return true
    }

    if A[x]!=A[y]{
        return false
    }

    x++
	y--
	return isPal(A, x, y)
}

```

##  Print 1 to A function

Problem Description

You are given an integer A, print 1 to A using using recursion.

Note :- After printing all the numbers from 1 to A, print a new line.



Problem Constraints

1 <= A <= 104



Input Format

First argument A is an integer.



Output Format

Print A space-separated integers 1 to A.
Note: There should be exactly one space after each integer. After printing all the integers, print a new line



Example Input

Input 1:

A = 10
Input 2:

A = 5


Example Output

Output 1:

1 2 3 4 5 6 7 8 9 10 
Output 2:

1 2 3 4 5 


Example Explanation

Explanation 1:

Print 1 to 10 space separated integers.
Explanation 2:

Print 1 to 5 space separated integers.

## Print reverse string

Problem Description

Write a recursive function that takes a string, S, as input and prints the characters of S in reverse order.



Problem Constraints

1 <= |s| <= 1000



Input Format

First line of input contains a string S.



Output Format

Print the character of the string S in reverse order.



Example Input

Input 1:

 scaleracademy
Input 2:

 cool


Example Output

Output 1:

 ymedacarelacs
Output 2:

 looc

 ## Sum of Digits!

 Problem Description

Given a number A, we need to find the sum of its digits using recursion.


Problem Constraints

1 <= A <= 109


Input Format

The first and only argument is an integer A.


Output Format

Return an integer denoting the sum of digits of the number A.


Example Input

Input 1:

 A = 46
Input 2:

 A = 11


Example Output

Output 1:

 10
Output 2:

 2


Example Explanation

Explanation 1:

 Sum of digits of 46 = 4 + 6 = 10
Explanation 2:

 Sum of digits of 11 = 1 + 1 = 2

 ``` go

 func solve(A int )  (int) {
    if A == 0 {
		return 0
	}

	return A%10 + solve(A/10)
}

```

## Print A to 1 function

Problem Description

You are given an integer A, print A to 1 using using recursion.

Note :- After printing all the numbers from A to 1, print a new line.



Problem Constraints

1 <= A <= 104



Input Format

First argument A is an integer.



Output Format

Print A space-separated integers A to 1.
Note: There should be exactly one space after each integer. Print a new line after printing the A integers



Example Input

Input 1:

10
Input 2:

5


Example Output

Output 1:

10 9 8 7 6 5 4 3 2 1 

Output 2:

5 4 3 2 1 


Example Explanation

Explanation 1:

Print 10 to 1 space separated integers.
Explanation 2:

Print 5 to 1 space separated integers.


## IS Magic

Problem Description

Given a number A, check if it is a magic number or not.

A number is said to be a magic number if the sum of its digits is calculated till a single digit recursively by adding the sum of the digits after every addition. If the single digit comes out to be 1, then the number is a magic number.



Problem Constraints

1 <= A <= 109



Input Format

The first and only argument is an integer A.



Output Format

Return an 1 if the given number is magic else return 0.



Example Input

Input 1:

 A = 83557
Input 2:

 A = 1291


Example Output

Output 1:

 1
Output 2:

 0


Example Explanation

Explanation 1:

 Sum of digits of (83557) = 28
 Sum of digits of (28) = 10
 Sum of digits of (10) = 1. 
 Single digit is 1, so it's a magic number. Return 1.
Explanation 2:

 Sum of digits of (1291) = 13
 Sum of digits of (13) = 4
 Single digit is not 1, so it's not a magic number. Return 0.


## Kth Symbol - easy 

Problem Description

On the first row, we write a 0. Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01, and each occurrence of 1 with 10.

Given row number A and index B, return the Bth indexed symbol in row A. (The values of B are 0-indexed.).



Problem Constraints

1 <= A <= 20

0 <= B < 2A - 1



Input Format

First argument is an integer A.

Second argument is an integer B.



Output Format

Return an integer denoting the Bth indexed symbol in row A.



Example Input

Input 1:

 A = 3
 B = 0
Input 2:

 A = 4
 B = 4


Example Output

Output 1:

 0
Output 2:

 1


Example Explanation

Explanation 1:

 Row 1: 0
 Row 2: 01
 Row 3: 0110
Explanation 2:

 Row 1: 0
 Row 2: 01
 Row 3: 0110
 Row 4: 01101001


# Kth Symmbol hard

Problem Description

On the first row, we write a 0. Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01, and each occurrence of 1 with 10.

Given row number A and index B, return the Bth indexed symbol in row A. (The values of B are 0-indexed.).



Problem Constraints

1 <= A <= 105

0 <= B <= min(2A - 1 - 1 , 1018)



Input Format

First argument is an integer A.

Second argument is an integer B.



Output Format

Return an integer denoting the Bth indexed symbol in row A.



Example Input

Input 1:

 A = 3
 B = 0
Input 2:

 A = 4
 B = 4


Example Output

Output 1:

 0
Output 2:

 1


Example Explanation

Explanation 1:

 Row 1: 0
 Row 2: 01
 Row 3: 0110
Explanation 2:

 Row 1: 0
 Row 2: 01
 Row 3: 0110
 Row 4: 01101001



Expected Output
Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases

``` go

func solve(A int , B int64 )  (int) {
    if B==0{
        return 0
    }
    val:=solve(A-1,B/2)
    if B%2==0{
        return val
    }
    return 1-val
}

```