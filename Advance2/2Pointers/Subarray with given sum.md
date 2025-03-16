## Subarray with given sum

```
Problem Description

Given an array of positive integers A and an integer B, find and return first continuous subarray which adds to B.






If the answer does not exist return an array with a single integer "-1".

First sub-array means the sub-array for which starting index in minimum.








Problem Constraints

1 <= length of the array <= 100000
1 <= A[i] <= 109
1 <= B <= 109



Input Format

The first argument given is the integer array A.

The second argument given is integer B.



Output Format

Return the first continuous sub-array which adds to B and if the answer does not exist return an array with a single integer "-1".



Example Input

Input 1:

 A = [1, 2, 3, 4, 5]
 B = 5
Input 2:

 A = [5, 10, 20, 100, 105]
 B = 110


Example Output

Output 1:





 [2, 3]
Output 2:

 [-1]






Example Explanation

Explanation 1:

 [2, 3] sums up to 5.
Explanation 2:

 No subarray sums up to required number.
```
``` go
/**
 * @input A : Integer array
 * @input B : Integer
 * 
 * @Output Integer array.
 */
func solve(A []int , sum int )  ([]int) {
    
   m := map[int]int{}
	
	ans := make([]int, 0)
	psum := 0
	s, e := -1, -1
	for i := 0; i < len(A); i++ {
		psum += A[i]
		if psum == sum {
			s = 0
			e = i
			break
		} else {
			if _, ok := m[psum-sum]; ok {
				s = m[psum-sum] + 1
				e = i
				break
			} else {
				m[psum] = i
			}
		}

	}

	
	if s == -1 && e == -1 {
		ans =append(ans,-1)
		
		return ans
	}
    for i := s; i <= e; i++ {
		ans = append(ans, A[i])
	}
	return ans
}
```
``` go
func FindSubarray(A []int, B int) []int {
    start, end := 0, 0
    currentSum := 0
    for end < len(A) {
        currentSum += A[end]
        for currentSum > B && start <= end {
            currentSum -= A[start]
            start++
        }
        if currentSum == B {
            return A[start : end+1]
        }
        end++
    }
    return []int{-1}
}
```