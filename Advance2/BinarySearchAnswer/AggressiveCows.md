## Aggressive Cows

```
Problem Description

Farmer John has built a new long barn with N stalls. Given an array of integers A of size N where each element of the array represents the location of the stall and an integer B which represents the number of cows.

His cows don't like this barn layout and become aggressive towards each other once put into a stall. To prevent the cows from hurting each other, John wants to assign the cows to the stalls, such that the minimum distance between any two of them is as large as possible. What is the largest minimum distance?



Problem Constraints

2 <= N <= 100000
0 <= A[i] <= 109
2 <= B <= N



Input Format

The first argument given is the integer array A.
The second argument given is the integer B.



Output Format

Return the largest minimum distance possible among the cows.



Example Input

Input 1:

A = [1, 2, 3, 4, 5]
B = 3
Input 2:

A = [1, 2]
B = 2


Example Output

Output 1:

 2
Output 2:

 1


Example Explanation

Explanation 1:

 John can assign the stalls at location 1, 3 and 5 to the 3 cows respectively. So the minimum distance will be 2.
Explanation 2:

 The minimum distance will be 1.



Expected Output
Provide sample input and click run to see the correct output for the provided input. Use this to improve your problem understanding and test edge cases

```

``` go
/**
 * @input A : Integer array
 * @input B : Integer
 * 
 * @Output Integer
 */
import "sort"
func solve(A []int , B int )  (int) {
    l:=1
    sort.Ints(A)
    r:=A[len(A)-1]-A[0]
   // var ans int
    for l<=r{
        mid:=l+(r-l)/2
        if check(A,B,mid){
            //ans=l
            l=mid+1
        }else{
            r=mid-1
        }
    }
    return r
}

func check(A []int, M int, mid int)bool{
    last_placed:=A[0]
    var cow int=1
    for i:=1;i<len(A);i++{
        if A[i]-last_placed>=mid{
            cow++
            last_placed=A[i]
            if cow==M{
                return true
            }
        }
    }
    return false
}


```