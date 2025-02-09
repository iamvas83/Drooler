package recirsion1

import "fmt"

type Result struct {
	res  [][]int
	curr int
}

func TowerOfHanoi(A int) [][]int {
	// Initialize the result struct
	result := &Result{}
	result.curr = 0
	l := (1 << A) - 1 // (2^A - 1) minimum moves required
	result.res = make([][]int, l)

	for x := range result.res {
		result.res[x] = make([]int, 3) // Each move stores [disk, source, destination]
	}

	solve(A, 1, 3, 2, result) // Start solving
	return result.res
}

func solve(A int, src int, dest int, temp int, result *Result) {
	if A == 0 {
		return
	}

	// Move (A-1) disks from src to temp using dest as auxiliary
	solve(A-1, src, temp, dest, result)

	// Store the move for the current disk
	result.res[result.curr] = []int{A, src, dest}
	result.curr++

	// Move (A-1) disks from temp to dest using src as auxiliary
	solve(A-1, temp, dest, src, result)
}

func main() {
	A := 3
	moves := TowerOfHanoi(A)
	fmt.Println("Moves:", moves)
}
