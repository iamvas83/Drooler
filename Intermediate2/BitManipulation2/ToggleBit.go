package bitmanipulation2

func ToggleIBit(A int, B int) int {
	mask := 1 << B
	res := A ^ mask
	return res
}
