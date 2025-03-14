## Painter's Partition Problem

```
Problem Description

Given 2 integers A and B and an array of integers C of size N. Element C[i] represents the length of ith board.
You have to paint all N boards [C0, C1, C2, C3 … CN-1]. There are A painters available and each of them takes B units of time to paint 1 unit of the board.

Calculate and return the minimum time required to paint all boards under the constraints that any painter will only paint contiguous sections of the board.
NOTE:
1. 2 painters cannot share a board to paint. That is to say, a board cannot be painted partially by one painter, and partially by another.
2. A painter will only paint contiguous boards. This means a configuration where painter 1 paints boards 1 and 3 but not 2 is invalid.

Return the ans % 10000003.



Problem Constraints

1 <= A <= 1000
1 <= B <= 106
1 <= N <= 105
1 <= C[i] <= 106



Input Format

The first argument given is the integer A.
The second argument given is the integer B.
The third argument given is the integer array C.



Output Format

Return minimum time required to paint all boards under the constraints that any painter will only paint contiguous sections of board % 10000003.



Example Input

Input 1:

 A = 2
 B = 5
 C = [1, 10]
Input 2:

 A = 10
 B = 1
 C = [1, 8, 11, 3]


Example Output

Output 1:

 50
Output 2:

 11

 ```
 ``` go
 import "math"
func paint(C int , B int , A []int )  (int) {
    var l int
    var r int
    max:=math.MinInt
    for i:=0;i<len(A);i++{
        if A[i]>max{
            max=A[i]

        }
    }
    var sum int
    for i:=0;i<len(A);i++{
        sum+=A[i]
    }
    l=max
    r=sum
    //var ans int
    for l<r{
        mid:=l+(r-l)/2

        if check(A,mid,C){
            r=mid
        }else{
            l=mid+1
        }
    }
    return (l*B)%10000003
}

func check(A []int, time int, C int)bool{
    w:=1
    var current_work int
    for i:=0;i<len(A);i++{
        
        if current_work+A[i]<=time{
            current_work+=A[i]
        }else{
            w++
            current_work=A[i]
        }
    }
    if w<=C{
        return true
    }else{
        return false
    }
}
```