## Kth Symbol - Hard

Problem Description

On the first row, we write a 0. Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01, and each occurrence of 1 with 10.

Given row number A and index B, return the Bth indexed symbol in row A. (The values of B are 0-indexed.).


```
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

 ```

 ``` go

 package main

import (
	"fmt"
)

// Recursive function to find the Kth symbol
func kthSymbol(A int, B int) int {
	if A == 1 {
		return 0
	}

	mid := 1 << (A - 2) // 2^(A-2) is the midpoint

	if B < mid {
		return kthSymbol(A-1, B) // Same as previous row
	} else {
		return 1 - kthSymbol(A-1, B-mid) // Flip the value
	}
}

func main() {
	fmt.Println(kthSymbol(3, 2))  // Output: 1
	fmt.Println(kthSymbol(4, 5))  // Output: 1
	fmt.Println(kthSymbol(10, 11)) // Output: 0
}

```

## Tower of Hanoi

Problem Description

In the classic problem of the Towers of Hanoi, you have 3 towers numbered from 1 to 3 (left to right) and A disks numbered from 1 to A (top to bottom) of different sizes which can slide onto any tower.
The puzzle starts with disks sorted in ascending order of size from top to bottom (i.e., each disk sits on top of an even larger one).
You have the following constraints:

Only one disk can be moved at a time.
A disk is slid off the top of one tower onto another tower.
A disk cannot be placed on top of a smaller disk.

You have to find the solution to the Tower of Hanoi problem.
You have to return a 2D array of dimensions M x 3, where M is the minimum number of moves needed to solve the problem.
In each row, there should be 3 integers (disk, start, end), where:

disk - number of the disk being moved
start - number of the tower from which the disk is being moved
end - number of the tower to which the disk is being moved



Problem Constraints

1 <= A <= 18


Input Format

The first argument is the integer A.


Output Format

Return a 2D array with dimensions M x 3 as mentioned above in the description.


Example Input

Input 1:
A = 2
Input 2:

A = 3


Example Output

Output 1:
[1 1 2 ] [2 1 3 ] [1 2 3 ]
Output 2:

[1 1 3 ] [2 1 2 ] [1 3 2 ] [3 1 3 ] [1 2 1 ] [2 2 3 ] [1 1 3 ] 


Example Explanation

Explanation 1:
We shift the first disk to the middle tower.
We shift the second disk to the last tower.
We, finally, shift the first disk from the middle tower to the last tower.
Explanation 2:

We can see that this was the only unique path with minimal moves to move all disks from the first to the third tower.

``` go

package main

import "fmt"

// Function to solve Tower of Hanoi
func towerOfHanoi(n, start, end, aux int, result *[][]int) {
    if n == 0 {
        return
    }

    // Move (n-1) disks from start to auxiliary peg
    towerOfHanoi(n-1, start, aux, end, result)

    // Move nth disk from start to end
    *result = append(*result, []int{n, start, end})

    // Move (n-1) disks from auxiliary to end peg
    towerOfHanoi(n-1, aux, end, start, result)
}

// Wrapper function
func solve(A int) [][]int {
    result := [][]int{}
    towerOfHanoi(A, 1, 3, 2, &result)
    return result
}

// Driver Code
func main() {
    A := 3
    moves := solve(A)
    fmt.Println("Moves:", moves)
}

```

## Is Magic or not?

```

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

```

``` go
func IsMagic(A int)int{
    if A%9==1{
        return 1
    }
    return 0
}
```