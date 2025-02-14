package advance2

import (
	"fmt"
)

func merge(arr []int, temp []int, left, mid, right int) int {
	i, j, k := left, mid+1, left
	invCount := 0

	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			invCount += (mid - i + 1) // Counting inversions
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}

	for x := left; x <= right; x++ {
		arr[x] = temp[x]
	}

	return invCount
}

func InversionCount_Sort(arr []int, temp []int, left, right int) int {
	invCount := 0
	if left < right {
		mid := (left + right) / 2
		invCount += InversionCount_Sort(arr, temp, left, mid)
		invCount += InversionCount_Sort(arr, temp, mid+1, right)
		invCount += merge(arr, temp, left, mid, right)
	}
	return invCount
}

func Main_inversion() {
	var n int
	fmt.Scan(&n) // Input size of the array

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i]) // Input array elements
	}

	temp := make([]int, n)
	invCount := InversionCount_Sort(arr, temp, 0, n-1)

	fmt.Println(invCount) // Output the inversion count
}
