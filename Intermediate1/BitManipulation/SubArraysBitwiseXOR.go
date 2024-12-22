package bitmanipulation

func SubArraysWithBitwiseXor(A []int) int64 {
	n := len(A)
	count := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			subSum := 0
			for k := i; k <= j; k++ {
				subSum |= A[k]
			}
			if subSum == 1 {
				count++
			}
		}

	}
	return int64(count)
}

func SubArraysWithBitwiseXor2(A []int) int64 {
	n := len(A)
	count := 0
	for i := 0; i < n; i++ {
		subSum := 0
		for j := i; j < n; j++ {

			subSum |= A[j]

			if subSum == 1 {
				count++
			}
		}

	}
	return int64(count)
}

/*
1. If Single '1' is present in Subarray, its bitwise OR=1
2. If all elements of a subarray are 0, its bitwise OR=0
Subarrays with Bitwise OR=1 = Total no. of Subarrays - total no. of subarrays with Bitwise OR=0
*/

func SubArraysWithBitwiseXor3(A []int) int64 {
	count := 0
	zeros := 0
	n := len(A)
	for i := 0; i < n; i++ {
		if A[i] == 0 {
			zeros++
		} else {
			count += (zeros * (zeros + 1)) / 2
			zeros = 0
		}
	}
	count += (zeros * (zeros + 1)) / 2
	return int64((n*(n+1))/2 - count)
}
