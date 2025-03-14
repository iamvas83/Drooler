## Special Integer

```
Problem Description

Given an array of integers A and an integer B, find and return the maximum value K such that there is no subarray in A of size K with the sum of elements greater than B.



Problem Constraints

1 <= |A| <= 100000
1 <= A[i] <= 10^9

1 <= B <= 10^9



Input Format

The first argument given is the integer array A.

The second argument given is integer B.



Output Format

Return the maximum value of K (sub array length).



Example Input

Input 1:

A = [1, 2, 3, 4, 5]
B = 10
Input 2:

A = [5, 17, 100, 11]
B = 130


Example Output

Output 1:

 2
Output 2:

 3
 ```
 ``` go
 /**
 * @input A : Integer array
 * @input B : Integer
 * 
 * @Output Integer
 */
func solve(A []int , B int )  (int) {
    n := len(A)

    // Helper function to check if a subarray of size `k` is valid
    isValid := func(k int) bool {
        sum := 0
        // Calculate the sum of the first `k` elements
        for i := 0; i < k; i++ {
            sum += A[i]
        }
        if sum > B {
            return false
        }
        // Sliding window to check other subarrays of size `k`
        for i := k; i < n; i++ {
            sum += A[i] - A[i-k]
            if sum > B {
                return false
            }
        }
        return true
    }

    // Binary search for the maximum size `K`
    left, right, result := 1, n, 0
    for left <= right {
        mid := (left + right) / 2
        if isValid(mid) {
            result = mid
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return result
}

```

