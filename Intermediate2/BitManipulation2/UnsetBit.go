package bitmanipulation2

/*
If B-th bit in A is set, make it unset.
If B-th bit in A is unset, leave as it is.
*/
func UnsetBit(A int, B int) int {
	mask := ^(1 << B)

	return A & mask
}
